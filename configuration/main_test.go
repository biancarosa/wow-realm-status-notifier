package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func generateFakeConfig() *Config {
	fakeConfig := new(Config)
	fakeConfig.TelegramToken = faker.Lorem().String()
	fakeConfig.Port = faker.Number().Number(5)
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

	assert.Panics(t, c.Validate, "The code did not panic when telegram token was not set")
}
