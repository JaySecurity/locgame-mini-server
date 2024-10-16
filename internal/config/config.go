package config

import (
	"locgame-mini-server/internal/config/resources"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"

	"github.com/kelseyhightower/envconfig"
)

// Environment indicates the current environment.
type Environment string

const (
	// Production points to the Production environment.
	Production Environment = "production"

	// Staging points to the Staging environment.
	Staging Environment = "staging"

	// Development points to the Development environment.
	Development Environment = "development"

	// Custom points to the Custom environment.
	Custom Environment = "custom"

	// Local points to the Local environment.
	Local Environment = "local"
)

type Config struct {
	*GameConfigs

	Environment Environment `default:"local"`

	HttpPort    int `default:"8080" split_words:"true"`
	MetricsPort int `default:"9090" split_words:"true"`

	NatsAddress string `default:"localhost:4222" split_words:"true"`

	RateLimit int `default:"500" split_words:"true"`

	// MinVersion - minimum client version x.x.x (For example 0.4.3 -> 43)
	MinVersion int `default:"10" split_words:"true"`

	Redis    *RedisConfig
	Database *DatabaseConfig
	Paypal   *PaypalConfig

	OverrideConfigBranch string `split_words:"true"`

	Repository Repository

	NetworkVerboseMode bool `default:"false" split_words:"true"`

	OnReload         func()
	OnReloadComplete []func()
}

// Init initializes configurations.
// Reads environment variables and loads matchmaking YAML config files.
// Accepts a relative path argument to the root of the config files.
// For example: "/", "", "../../".
func Init(path string) *Config {
	log.Debug("Path: " + path)
	parentPath = path
	c := new(Config)
	if err := envconfig.Process("", c); err != nil {
		panic(err)
	}

	branch := c.OverrideConfigBranch
	log.Debug("Branch: " + branch)
	if branch == "" {
		branch = string(c.Environment)
	}

	c.Repository = New(path, branch)
	c.GameConfigs = new(GameConfigs)
	c.Paypal = NewPaypalConfig()
	c.Paypal.SetEnvironment(c.Environment)
	c.load()
	pubsub.RegisterHandler(&ConfigsReloadHandler{config: c})

	return c
}

func (c *Config) load() {
	resources.Init(parentPath)

	c.GameConfigs.Load()

	if err := envconfig.Process("", c.GameConfigs); err != nil {
		panic(err)
	}

	c.Blockchain.SetEnvironment(c.Environment)
	c.Ses.SetEnvironment(c.Environment)
}

func (c *Config) Reload() {
	if c.OnReload != nil {
		c.OnReload()
	}

	if c.Repository.RefreshWorkspace() {
		c.load()

		log.Info("Configs successfully reloaded.")
	}

	if len(c.OnReloadComplete) > 0 {
		for _, callback := range c.OnReloadComplete {
			callback()
		}
	}
}
