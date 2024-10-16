package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewCognitoConfig)
}

// CognitoConfig stores cognito configuration.
type CognitoConfig struct {
	BaseConfig

	REGION       string `yaml:"REGION"`
	REDIRECT_URI string `yaml:"REDIRECT_URI"`

	AWS_EMAIL_CLIENT_ID string `yaml:"AWS_EMAIL_CLIENT_ID"`
	EMAIL_DOMAIN        string `yaml:"EMAIL_DOMAIN"`
	EMAIL_USER_POOL_ID  string `yaml:"EMAIL_USER_POOL_ID"`
	EMAIL_CLIENT_SECRET string `yaml:"EMAIL_CLIENT_SECRET"`

	AWS_SOCIAL_CLIENT_ID    string `yaml:"AWS_SOCIAL_CLIENT_ID"`
	SOCIAL_USER_POOL_ID     string `yaml:"SOCIAL_USER_POOL_ID"`
	SOCIAL_CLIENT_SECRET    string `yaml:"SOCIAL_CLIENT_SECRET"`
	SOCIAL_IDENTITY_POOL_ID string `yaml:"SOCIAL_IDENTITY_POOL_ID"`
}

// NewCognitoConfig creates an instance of the cognito configuration.
func NewCognitoConfig() *CognitoConfig {
	c := new(CognitoConfig)
	c.self = c
	c.Load("cognito.yaml")
	return c
}

func (c *CognitoConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		if err = yaml.Unmarshal(bytes, &c); err != nil {
			return err
		}
	}

	return nil
}
