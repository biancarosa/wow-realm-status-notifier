package configuration

import (
	"errors"
)

var c *Config

type Config struct {
	TelegramToken   string
	Port            string
	GoogleProjectID string
}

func (c *Config) Validate() error {
	if c.Port == "" {
		c.Port = "3000"
	}
	if c.TelegramToken == "" {
		return errors.New("telegram token must be provided")
	}
	return nil
}

func GetConfig() *Config {
	if c != nil {
		return c
	}
	c = new(Config)
	ep := new(EnvironmentVariablesConfigurationProvider)
	c = ep.PopulateConfig(c)
	if err := c.Validate(); err == nil {
		return c
	}

	gsmp := new(GoogleSecretManagerConfigurationProvider)
	c = gsmp.PopulateConfig(c)
	if err := c.Validate(); err != nil {
		panic(err)
	}
	return c
}
