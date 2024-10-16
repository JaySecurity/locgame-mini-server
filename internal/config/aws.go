package config

func init() {
	Register(NewAwsConfig)
}

// AwsConfig stores aws configuration.
type AwsConfig struct {
	AccessKeyID     string `default:"USE_ONLY_ON_PROD_AND_DEV_ENV" required:"true" split_words:"true"`
	SecretAccessKey string `default:"USE_ONLY_ON_PROD_AND_DEV_ENV" required:"true" split_words:"true"`
	Region          string `default:"us-east-2" split_words:"true"`
}

// NewAwsConfig creates an instance of the aws configuration.
func NewAwsConfig() *AwsConfig {
	c := new(AwsConfig)
	return c
}
