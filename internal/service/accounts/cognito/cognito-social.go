package cognito

import (
	"context"
	"fmt"
	envConfig "locgame-mini-server/internal/config"
	common "locgame-mini-server/internal/service/accounts/lib"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"
)

type Cognito struct {
	identityClient *cognitoidentity.Client
	providerClient *cognitoidentityprovider.Client

	env *envConfig.CognitoConfig
}

func NewCognito(env *envConfig.CognitoConfig) *Cognito {
	//Load AWS configuration

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(env.REGION))
	if err != nil {
		log.Fatal("NewCognito: ", err)
	}
	providerClient := cognitoidentityprovider.NewFromConfig(cfg)
	identityClient := cognitoidentity.NewFromConfig(cfg)
	return &Cognito{
		providerClient: providerClient,
		identityClient: identityClient,
		env:            env,
	}
}

func (c *Cognito) getUserFromCognito(ctx context.Context, username string) error {
	params := &cognitoidentityprovider.AdminGetUserInput{
		Username:   aws.String(username),
		UserPoolId: aws.String(c.env.SOCIAL_USER_POOL_ID),
	}
	_, err := c.providerClient.AdminGetUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
func (c *Cognito) createCognitoUser(u *common.UserData) error {
	ctx := context.TODO()
	params := &cognitoidentityprovider.AdminCreateUserInput{
		Username:   aws.String(u.User.Email),
		UserPoolId: aws.String(c.env.SOCIAL_USER_POOL_ID),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(u.User.Email)},
			{Name: aws.String("email_verified"), Value: aws.String(fmt.Sprintf("%t", u.User.Email_verified))},
			{Name: aws.String("family_name"), Value: aws.String(u.User.Last_name)},
			{Name: aws.String("given_name"), Value: aws.String(u.User.First_name)},
		},
		MessageAction:     types.MessageActionTypeSuppress,
		TemporaryPassword: aws.String(common.CreateHash(u.User.Username, u.Provider, c.env.AWS_SOCIAL_CLIENT_ID)),
	}
	user, err := c.providerClient.AdminCreateUser(ctx, params)
	if err != nil {
		return err
	}
	u.User.Username = *user.User.Username

	setPass := &cognitoidentityprovider.AdminSetUserPasswordInput{
		Password:   aws.String(common.CreateHash(u.User.Username, u.Provider, c.env.AWS_SOCIAL_CLIENT_ID)),
		Permanent:  true,
		UserPoolId: aws.String(c.env.SOCIAL_USER_POOL_ID),
		Username:   aws.String(u.User.Username),
	}
	_, err = c.providerClient.AdminSetUserPassword(ctx, setPass)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cognito) getIdentityId(idToken string) (string, error) {

	input := &cognitoidentity.GetIdInput{
		IdentityPoolId: &c.env.SOCIAL_IDENTITY_POOL_ID,
		Logins: map[string]string{
			"accounts.google.com": idToken,
		},
	}

	output, err := c.identityClient.GetId(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to get id, %v", err)
	}
	if output.IdentityId == nil {
		return "", fmt.Errorf("failed to get identity id")
	}
	return *output.IdentityId, nil
}

func (c *Cognito) authenticateUser(idToken string, u *common.UserData) error {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeCustomAuth,
		AuthParameters: map[string]string{
			"USERNAME":    u.User.Email,
			"SECRET_HASH": common.CreateHash(c.env.SOCIAL_CLIENT_SECRET, c.env.AWS_SOCIAL_CLIENT_ID, u.User.Email),
		},
		ClientId: &c.env.AWS_SOCIAL_CLIENT_ID,
		// UserPoolId: aws.String(c.UserPoolId),
		ClientMetadata: map[string]string{
			"idToken":  idToken,
			"provider": u.Provider,
		},
	}
	authOutput, err := c.providerClient.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		if strings.Contains(err.Error(), "NotAuthorizedException: Incorrect username or password.") {
			//create User
			return fmt.Errorf("unable to authenticate user")
		} else {
			log.Panicf("failed to get ID, %v", err)
		}
	}
	u.User.Username = authOutput.ChallengeParameters["USERNAME"]
	challengeResponseInput := &cognitoidentityprovider.RespondToAuthChallengeInput{
		ChallengeName: types.ChallengeNameTypeCustomChallenge,
		ClientId:      &c.env.AWS_SOCIAL_CLIENT_ID,
		// UserPoolId:    aws.String(c.UserPoolId),
		ChallengeResponses: map[string]string{
			"USERNAME":    u.User.Username,
			"SECRET_HASH": common.CreateHash(c.env.SOCIAL_CLIENT_SECRET, c.env.AWS_SOCIAL_CLIENT_ID, u.User.Username),
			"ANSWER":      common.CreateHash(u.User.Username, u.Provider, c.env.AWS_SOCIAL_CLIENT_ID),
		},
		ClientMetadata: map[string]string{
			"provider": u.Provider,
		},
		Session: authOutput.Session,
	}
	respOutput, err := c.providerClient.RespondToAuthChallenge(context.TODO(), challengeResponseInput)
	if err != nil {
		log.Panicf("failed to get Token, %v", err)
	}
	if respOutput.AuthenticationResult != nil {
		u.IdToken = *respOutput.AuthenticationResult.IdToken
		u.RefreshToken = *respOutput.AuthenticationResult.RefreshToken
		u.AccessToken = *respOutput.AuthenticationResult.AccessToken
		return nil
	}
	return fmt.Errorf("failed to Authenticate User")
}

func (c *Cognito) HandleSocialAuth(l *common.LoginRequest, u *common.UserData) error {

	identityId, err := c.getIdentityId(l.IdToken)
	if err != nil {
		return err
	}

	u.User.IdentityId = identityId

	err = c.getUserFromCognito(context.TODO(), u.User.Email)
	if err != nil {
		if strings.Contains(err.Error(), "UserNotFoundException") {
			//create User
			err = c.createCognitoUser(u)
			if err != nil {
				return fmt.Errorf("unable to create user: %v", err)
			}
		} else {
			return fmt.Errorf("failed to find Cognito user, %v", err)
		}
	}

	err = c.authenticateUser(l.IdToken, u)
	if err != nil {
		return err
	}

	return nil
}
