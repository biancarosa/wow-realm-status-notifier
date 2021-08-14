package database

import (
	"github.com/biancarosa/wow-realm-status-notifier/configuration"
	"github.com/biancarosa/wow-realm-status-notifier/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	config := configuration.GetConfig()
	var err error
	if config.SQLLite != nil {
		DB, err = gorm.Open(sqlite.Open(config.SQLLite.Name), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	// Migrate the schema
	DB.AutoMigrate(&models.NotificationsRequest{})
}
