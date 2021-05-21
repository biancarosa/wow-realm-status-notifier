package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerFromTextWithBadInput(t *testing.T) {
	s, err := getServerFromText("badinput")
	assert.NotNil(t, err)
	assert.Empty(t, s)
}

func TestGetServerFromText(t *testing.T) {
	s, err := getServerFromText("good input")
	assert.Nil(t, err)
	assert.Equal(t, s, "input")
}
