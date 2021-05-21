package notifications_request

import "github.com/biancarosa/wow-realm-status-notifier/models"

func New(chatID int64, server string) *models.NotificationsRequest {
	n := new(models.NotificationsRequest)
	n.ChatID = chatID
	n.Server = server
	// TODO: Add current datetime
	return n
}
