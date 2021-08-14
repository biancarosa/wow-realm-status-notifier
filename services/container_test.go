package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	servicesContainer := GetServices(context.Background())
	assert.NotNil(t, servicesContainer)
}
