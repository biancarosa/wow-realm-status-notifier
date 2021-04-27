package configuration

import "os"

type ConfigurationProvider interface {
	GetConfig() *Config
}

type EnvironmentVariablesConfigurationProvider struct{}

func (p *EnvironmentVariablesConfigurationProvider) GetConfig() *Config {
	c := new(Config)
	c.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	c.Port = os.Getenv("PORT")
	return c
}
