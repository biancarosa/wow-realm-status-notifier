package services

import (
	"context"

	"github.com/biancarosa/wow-realm-status-notifier/services/notifications_request_service"
)

type AppServices struct {
	NotificationsRequestService notifications_request_service.NotificationReqServiceInterface
}

func GetServices(ctx context.Context) *AppServices {
	return &AppServices{
		NotificationsRequestService: notifications_request_service.New(ctx),
	}
}
