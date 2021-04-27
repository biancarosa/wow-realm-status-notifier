package configuration

type Config struct {
	TelegramToken string
	Port          string
}

func (c *Config) Validate() {
	if c.Port == "" {
		c.Port = "3000"
	}
	if c.TelegramToken == "" {
		panic("Telegram Token must be provided")
	}
}

func GetConfig() *Config {
	provider := new(EnvironmentVariablesConfigurationProvider)
	c := provider.GetConfig()
	c.Validate()
	return c
}
