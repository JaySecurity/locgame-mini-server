package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	IdToken string `json:"id_token"`
}

type UserData struct {
	IdToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Provider     string `json:"iss"`
	Sub          string `json:"sub"`
	User         struct {
		Username       string `json:"username" bson:"username"`
		First_name     string `json:"given_name" bson:"first_name"`
		Last_name      string `json:"family_name" bson:"last_name"`
		Email_verified bool   `json:"email_verified" bson:"email_verified"`
		Email          string `json:"email" bson:"email"`
		IdentityId     string `json:"identityId" bson:"identityId"`
	} `json:"user" bson:"user"`
}

func ProcessMessage(idToken string, lr *LoginRequest, u *UserData) {
	// Parse the stringified JSON into a struct or map

	lr.IdToken = idToken
	parser := &jwt.Parser{
		SkipClaimsValidation: true,
	}
	decoded, _, err := parser.ParseUnverified(lr.IdToken, jwt.MapClaims{})

	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims.")
		return
	}

	u.Provider = claims["iss"].(string)
	u.Sub = claims["sub"].(string)

	u.User.First_name = claims["given_name"].(string)
	u.User.Last_name = claims["family_name"].(string)
	u.User.Email_verified = claims["email_verified"].(bool)
	u.User.Email = claims["email"].(string)
}

func CreateHash(hashKey string, partOne string, partTwo string) string {
	message := []byte(partTwo + partOne)
	key := []byte(hashKey)
	hasher := hmac.New(sha256.New, key)
	hasher.Write(message)
	hash := hasher.Sum(nil)
	secretHash := base64.StdEncoding.EncodeToString(hash)
	return secretHash
}

func ValidateEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Match the email against the pattern
	matched := regex.MatchString(email)

	return matched
}

func ExtractUsername(email string) string {
	// Split the email at the @ symbol
	parts := strings.Split(email, "@")

	// Check if parts are extracted correctly
	if len(parts) != 2 {
		return ""
	}

	// Return the username
	return parts[0]
}

func GenerateSignupInput(email, password, clientId string) *cognitoidentityprovider.SignUpInput {
	return &cognitoidentityprovider.SignUpInput{
		Username: aws.String(email),
		Password: aws.String(password),
		ClientId: aws.String(clientId),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("custom:terms"),
				Value: aws.String("agreed"),
			},
			{
				Name:  aws.String("given_name"),
				Value: aws.String(ExtractUsername(email)),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(ExtractUsername(email)),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(" "),
			},
		},
	}
}
