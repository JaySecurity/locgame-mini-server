package awsses

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	envConfig "locgame-mini-server/internal/config"
)

type SES struct {
	sesClient *ses.Client
	env       *envConfig.SesConfig
}

func NewSES(env *envConfig.SesConfig) *SES {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(env.REGION))
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil
	}
	return &SES{
		sesClient: ses.NewFromConfig(cfg),
		env:       env,
	}
}

func (s *SES) SendEmail(to, subject, htmlBody string, textBody string) error {

	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(htmlBody),
				},
				Text: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(textBody),
				},
			},
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.env.SENDER_ADDRESS),
	}

	_, err := s.sesClient.SendEmail(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
