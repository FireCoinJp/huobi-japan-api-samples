package config

import (
	"os"

	"go.uber.org/config"
)

type (
	Config struct {
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
		AccountID string `yaml:"account_id"`
		Host      string `yaml:"host"`
		Save      bool   `yaml:"save"`
	}
)

func Load(file string) (*Config, error) {
	reader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	provider, err := config.NewYAML(config.Source(reader))
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := provider.Get("setting").Populate(c); err != nil {
		return nil, err
	}
	return c, nil
}

var Cfg *Config
