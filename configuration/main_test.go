package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func generateFakeConfig() *Config {
	fakeConfig := new(Config)
	fakeConfig.TelegramToken = faker.Lorem().String()
	fakeConfig.Port = faker.Number().Number(5)
	fakeConfig.GoogleProjectID = faker.Lorem().String()
	return fakeConfig
}

func TestValidateWhenPortIsSet(t *testing.T) {
	c := generateFakeConfig()
	c.Port = "5000"

	c.Validate()

	assert.Equal(t, c.Port, "5000")
}

func TestValidateWhenPortIsNotSet(t *testing.T) {
	c := generateFakeConfig()
	c.Port = ""

	c.Validate()

	assert.Equal(t, c.Port, "3000")
}
func TestValidateWhenTelegramTokenIsNotSet(t *testing.T) {
	c := generateFakeConfig()
	c.TelegramToken = ""

	err := c.Validate()
	assert.NotNil(t, err)
}

func TestGetConfig(t *testing.T) {
	fakeConfig := generateFakeConfig()

	os.Setenv("TELEGRAM_TOKEN", fakeConfig.TelegramToken)
	os.Setenv("PORT", fakeConfig.Port)
	c := GetConfig()

	assert.Equal(t, c.TelegramToken, fakeConfig.TelegramToken)
	assert.Equal(t, c.Port, fakeConfig.Port)
}
