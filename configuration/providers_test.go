package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigFromEnvVars(t *testing.T) {
	fakeConfig := generateFakeConfig()

	os.Setenv("TELEGRAM_TOKEN", fakeConfig.TelegramToken)
	os.Setenv("PORT", fakeConfig.Port)

	p := new(EnvironmentVariablesConfigurationProvider)
	c := p.GetConfig()

	assert.Equal(t, c.TelegramToken, fakeConfig.TelegramToken)
	assert.Equal(t, c.Port, fakeConfig.Port)
}
