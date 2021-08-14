package notifications_request_service

import (
	"context"
	"os"
	"testing"

	"github.com/biancarosa/wow-realm-status-notifier/database"
	"github.com/biancarosa/wow-realm-status-notifier/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

func TestCreate(t *testing.T) {
	database.DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	database.DB.AutoMigrate(&models.NotificationsRequest{})
	s := New(context.Background())
	chatID := uint8(faker.RandomInt(1, 1000))
	server := faker.Lorem().String()
	nr := models.NotificationsRequest{
		ChatID: chatID,
		Server: server,
	}
	err := s.Create(&nr)
	assert.Nil(t, err)
	os.Remove("test.db")
}
