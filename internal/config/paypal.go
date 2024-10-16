package config

import (
	"locgame-mini-server/pkg/log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewPaypalConfig)
}

type PaypalEnv struct {
	URL       string `yaml:"BaseUrl"`
	ApproveId string `yaml:"ApproveId"`
	CaptureId string `yaml:"CaptureId"`
}

type PaypalConfig struct {
	BaseConfig
	*PaypalEnv
	environments map[string]*PaypalEnv `ignore:"true"`

	ClientId string `default:"" split_words:"true"`
	Secret   string `default:"" split_words:"true"`
}

func NewPaypalConfig() *PaypalConfig {
	c := new(PaypalConfig)
	c.self = c
	c.Load("paypal.yaml")
	if err := envconfig.Process("PAYPAL", c); err != nil {
		panic(err)
	}
	return c
}

func (c *PaypalConfig) Unmarshal() error {
	data := struct {
		Environments map[string]*PaypalEnv `yaml:"Environments"`
	}{}
	bytes, err := os.ReadFile(c.filePath)
	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err != nil {
			return err
		}
	}
	c.environments = data.Environments
	return nil
}

func (c *PaypalConfig) SetEnvironment(environment Environment) {
	caser := cases.Title(language.English)
	envName := caser.String(string(environment))
	if env, ok := c.environments[envName]; ok {
		c.PaypalEnv = env
		log.Debug("Selected blockchain environment:", envName)
	} else {
		log.Warning("Invalid blockchain environment:", envName, "Development environment will be used.")
		c.PaypalEnv = c.environments["Development"]
	}
}
