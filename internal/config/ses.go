package config

import (
	"locgame-mini-server/pkg/log"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewSesConfig)
}

// SesConfig stores AWS SES configuration.
type SesEnv struct {
	REGION         string `yaml:"REGION"`
	SENDER_ADDRESS string `yaml:"SENDER_ADDRESS"`
}
type SesConfig struct {
	BaseConfig
	*SesEnv

	environments map[string]*SesEnv `ignore:"true"`
}

// NewCognitoConfig creates an instance of the cognito configuration.
func NewSesConfig() *SesConfig {
	c := new(SesConfig)
	c.self = c
	c.Load("ses.yaml")
	return c
}

func (c *SesConfig) Unmarshal() error {
	bytes, err := os.ReadFile(c.filePath)
	data := struct {
		Environments map[string]*SesEnv `yaml:"Environments"`
	}{}

	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err != nil {
			return err
		}
	}
	c.environments = data.Environments
	return nil
}

func (c *SesConfig) SetEnvironment(environment Environment) {
	caser := cases.Title(language.English)
	envName := caser.String(string(environment))
	if env, ok := c.environments[envName]; ok {
		c.SesEnv = env
		log.Debug("Selected blockchain environment:", envName)
	} else {
		log.Warning("Invalid blockchain environment:", envName, "Development environment will be used.")
		c.SesEnv = c.environments["Development"]
	}
}
