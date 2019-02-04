package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API string `envconfig:"api_uri"`
}

func New() (*Config, error) {
	var c Config
	err := envconfig.Process(``, &c)
	return &c, err
}
