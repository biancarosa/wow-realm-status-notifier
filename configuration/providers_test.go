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
	os.Setenv("GOOGLE_PROJECT_ID", fakeConfig.GoogleProjectID)

	p := new(EnvironmentVariablesConfigurationProvider)
	c := p.PopulateConfig(new(Config))

	assert.Equal(t, c.TelegramToken, fakeConfig.TelegramToken)
	assert.Equal(t, c.Port, fakeConfig.Port)
	assert.Equal(t, c.GoogleProjectID, fakeConfig.GoogleProjectID)
}
