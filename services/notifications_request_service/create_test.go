package notifications_request_service

import (
	"testing"

	"github.com/biancarosa/wow-realm-status-notifier/models"
	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func TestCreate(t *testing.T) {
	s := New()
	chatID := uint8(faker.RandomInt(1, 1000))
	server := faker.Lorem().String()
	nr := models.NotificationsRequest{
		ChatID: chatID,
		Server: server,
	}
	err := s.Create(&nr)
	assert.NotNil(t, err)
}
