package database

import (
	"github.com/biancarosa/wow-realm-status-notifier/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dbName string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.NotificationsRequest{})
}
