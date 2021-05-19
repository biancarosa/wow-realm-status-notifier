package notifications_request

import "github.com/biancarosa/wow-realm-status-notifier/models"

func New(user string, server string) *models.NotificationsRequest {
	n := new(models.NotificationsRequest)
	n.User = user
	n.Server = server
	// TODO: Add current datetime
	return n
}
