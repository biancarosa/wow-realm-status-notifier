package notifications_request_service

import (
	"fmt"

	"github.com/biancarosa/wow-realm-status-notifier/database"
	"github.com/biancarosa/wow-realm-status-notifier/models"
)

func (s *notificationsRequestService) Create(nr *models.NotificationsRequest) error {
	tx := database.DB.Create(nr)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}
	return nil
}
