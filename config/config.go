package config

import (
  "github.com/kelseyhightower/envconfig"
)

type Config struct {
  ClientID  string `split_words:"true"`
  ClientSecret string `split_words:"true"`
}

func LoadConfig () (*Config, error) {
  c := &Config{}
  err := envconfig.Process("seaside", c)
  envconfig.Usage("seaside", c)
  if err != nil {
    return nil,err
  }
  return c, nil
}
