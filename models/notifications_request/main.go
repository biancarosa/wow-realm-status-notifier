package notifications_request

import "github.com/biancarosa/wow-realm-status-notifier/models"

func New(chatID uint8, server string) *models.NotificationsRequest {
	n := new(models.NotificationsRequest)
	n.ChatID = chatID
	n.Server = server
	return n
}
