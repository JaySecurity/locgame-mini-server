package cognito

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"
	common "locgame-mini-server/internal/service/accounts/lib"
	"locgame-mini-server/pkg/log"
)

func (c *Cognito) GetUserByEmail(ctx context.Context, username string) error {
	params := &cognitoidentityprovider.AdminGetUserInput{
		Username:   aws.String(username),
		UserPoolId: aws.String(c.env.EMAIL_USER_POOL_ID),
	}
	_, err := c.providerClient.AdminGetUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cognito) Signup(email string) error {
	password := common.CreateHash(c.env.EMAIL_CLIENT_SECRET, c.env.AWS_EMAIL_CLIENT_ID, email)
	user := common.GenerateSignupInput(email, password, c.env.AWS_EMAIL_CLIENT_ID)
	_, err := c.providerClient.SignUp(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cognito) SendEmail(email string) (string, string) {
	startAuthResp, err := c.providerClient.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(c.env.AWS_EMAIL_CLIENT_ID),
		AuthFlow: "CUSTOM_AUTH", // Specify custom auth flow
		AuthParameters: map[string]string{
			"USERNAME": email,
		},
	})
	if err != nil {
		log.Error("SendEmail", err)
	}
	return fmt.Sprint(startAuthResp.ChallengeName), *startAuthResp.Session
}

func (c *Cognito) VerifyEmail(email, code, challenge, session string) (string, string, error) {
	respondAuthResp, err := c.providerClient.RespondToAuthChallenge(context.TODO(), &cognitoidentityprovider.RespondToAuthChallengeInput{
		ClientId:      aws.String(c.env.AWS_EMAIL_CLIENT_ID),
		ChallengeName: types.ChallengeNameType(challenge),
		Session:       aws.String(session),
		ChallengeResponses: map[string]string{
			"USERNAME": email,
			"ANSWER":   code,
		},
	})
	if err != nil {
		log.Error("VerifyEmail: ", err)
		return "", "", err
	}

	return *respondAuthResp.AuthenticationResult.AccessToken, *respondAuthResp.AuthenticationResult.IdToken, nil
}

func (c *Cognito) GetUserByToken(ctx context.Context, accessToken string) (string, error) {
	params := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	}
	resp, err := c.providerClient.GetUser(ctx, params)
	if err != nil {
		return "", err
	}

	return *resp.Username, nil
}
