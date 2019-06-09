package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
  ConfigPrefix = "seaside"
  // Version FIXME: 本当にここに置くべきか?
  Version = "v0.0.1"
)

type Config struct {
	SeaOrigin      string `split_words:"true" default:"https://c.linkage.community"`
	ClientID       string `split_words:"true" required:"true"`
	ClientSecret   string `split_words:"true" required:"true"`
	CredentialFile string `split_words:"true" default:"credential.json"`
}

func Usage() {
	c := &Config{}
	envconfig.Usage(ConfigPrefix, c)
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	err := envconfig.Process(ConfigPrefix, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
