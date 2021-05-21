package notifications_request

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func TestNew(t *testing.T) {
	chatID := int64(faker.RandomInt(1, 1000))
	server := faker.Lorem().String()
	nr := New(chatID, server)
	assert.NotNil(t, nr)
	assert.Equal(t, nr.ChatID, chatID)
	assert.Equal(t, nr.Server, server)
}
