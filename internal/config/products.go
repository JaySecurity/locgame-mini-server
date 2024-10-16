package config

import (
	"io/ioutil"

	"locgame-mini-server/pkg/dto/store"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewProducts)
}

// Products stores products configuration.
type Products struct {
	BaseConfig

	ProductsByID map[string]Product
}

// NewProducts creates an instance of the products configuration.
func NewProducts() *Products {
	c := new(Products)
	c.self = c
	c.Load("store/products.yaml")
	return c
}

type Product struct {
	Type       store.ProductType `yaml:"Type"`
	Value      string            `yaml:"Value"`
	PriceInUSD float32           `yaml:"PriceInUSD"`
	PriceInLC  float32           `yaml:"PriceInLC"`
	Quantity   int64             `yaml:"Quantity"`
}

func (c *Products) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		err = yaml.Unmarshal(bytes, &c.ProductsByID)
	}

	return err
}
