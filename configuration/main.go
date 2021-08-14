package configuration

import (
	"errors"
)

var c *Config

type Config struct {
	TelegramToken   string
	Port            string
	GoogleProjectID string
	// MongoDB         *mongoDBConfig
	SQLLite *sqlLiteConfig
}

type sqlLiteConfig struct {
	Name string
}

// type mongoDBConfig struct {
// 	Host       string
// 	Port       string
// 	Database   string
// 	Collection string
// }

func (c *Config) Validate() error {
	if c.Port == "" {
		c.Port = "3000"
	}
	if c.TelegramToken == "" {
		return errors.New("telegram token must be provided")
	}
	// if c.MongoDB.Host == "" {
	// 	c.MongoDB.Host = "localhost"
	// }
	// if c.MongoDB.Port == "" {
	// 	c.MongoDB.Port = "27017"
	// }
	// if c.MongoDB.Database == "" {
	// 	c.MongoDB.Database = "notification_requests"
	// }
	// if c.MongoDB.Collection == "" {
	// 	c.MongoDB.Collection = "servers"
	// }
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
