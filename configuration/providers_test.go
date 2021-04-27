package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func TestGetConfigFromEnvVars(t *testing.T) {
	fakeConfig := new(Config)
	fakeConfig.TelegramToken = faker.Lorem().String()
	fakeConfig.Port = faker.Number().Number(5)
	os.Setenv("TELEGRAM_TOKEN", fakeConfig.TelegramToken)
	os.Setenv("PORT", fakeConfig.Port)
	p := new(EnvironmentVariablesConfigurationProvider)
	c := p.GetConfig()
	assert.Equal(t, c.TelegramToken, fakeConfig.TelegramToken)
	assert.Equal(t, c.Port, fakeConfig.Port)
}
