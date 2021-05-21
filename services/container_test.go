package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	servicesContainer := GetServices()
	assert.NotNil(t, servicesContainer)
}
