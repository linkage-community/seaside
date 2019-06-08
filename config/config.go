package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ClientID       string `split_words:"true" required:"true"`
	ClientSecret   string `split_words:"true" required:"true"`
	CredentialFile string `split_words:"true" default:"credential.json"`
}

func Usage() {
	c := &Config{}
	envconfig.Usage("seaside", c)
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("seaside", c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
