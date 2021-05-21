package models

import (
	"gorm.io/gorm"
)

type NotificationsRequest struct {
	gorm.Model
	Server string
	ChatID uint8
}
