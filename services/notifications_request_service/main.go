package notifications_request_service

import (
	"github.com/biancarosa/wow-realm-status-notifier/models"
)

type notificationReqServiceInterface interface {
	Create(*models.NotificationsRequest) error
}

type notificationsRequestService struct {
}

func New() *notificationsRequestService {
	s := new(notificationsRequestService)
	return s
}
