package config

import (
	"io/ioutil"

	"locgame-mini-server/pkg/log"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewBlockchain)
}

type BlockchainEnvironment struct {
	Contracts struct {
		NFT      string `yaml:"NFT"`
		USDT     string `yaml:"USDT"`
		USDC     string `yaml:"USDC"`
		LOCG     string `yaml:"LOCG"`
		BaseLOCG string `yaml:"BaseLOCG"`
		BaseUSDC string `yaml:"BaseUSDC"`
	} `yaml:"Contracts"`
	PaymentRecipients struct {
		USDT string `yaml:"USDT"`
		USDC string `yaml:"USDC"`
		LOCG string `yaml:"LOCG"`
	} `yaml:"PaymentRecipients"`
	RpcAddresses struct {
		Polygon  string `yaml:"Polygon"`
		Ethereum string `yaml:"Ethereum"`
		Base     string `yaml:"Base"`
	} `yaml:"RpcAddresses"`
	ChainNames struct {
		Polygon  string `yaml:"Polygon"`
		Ethereum string `yaml:"Ethereum"`
		Base     string `yaml:"Base"`
	} `yaml:"ChainNames"`
	ChainIds struct {
		Polygon  string `yaml:"Polygon"`
		Ethereum string `yaml:"Ethereum"`
		Base     string `yaml:"Base"`
	} `yaml:"ChainIds"`
}

// Blockchain stores blockchain configuration.
type Blockchain struct {
	BaseConfig
	*BlockchainEnvironment

	environments        map[string]*BlockchainEnvironment `ignore:"true"`
	MinterPrivateKey    string                            `default:"USE_ONLY_ON_PROD_AND_DEV_ENV" required:"true" split_words:"true"`
	CoinMarketCapApiKey string                            `default:"a2b38142-3639-4a33-8106-90ec65695348" required:"true" split_words:"true"`
}

// NewBlockchain creates an instance of the blockchain configuration.
func NewBlockchain() *Blockchain {
	c := new(Blockchain)
	c.self = c
	c.Load("blockchain.yaml")

	return c
}

func (c *Blockchain) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	var data struct {
		Environments map[string]*BlockchainEnvironment `yaml:"Environments"`
	}

	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err != nil {
			return err
		}
	}

	c.environments = data.Environments

	return nil
}

func (c *Blockchain) SetEnvironment(environment Environment) {
	caser := cases.Title(language.English)
	envName := caser.String(string(environment))
	if env, ok := c.environments[envName]; ok {
		c.BlockchainEnvironment = env
		log.Debug("Selected blockchain environment:", envName)
	} else {
		log.Warning("Invalid blockchain environment:", envName, "Development environment will be used.")
		c.BlockchainEnvironment = c.environments["Development"]
	}
}
