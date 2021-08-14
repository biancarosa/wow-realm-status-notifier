package notifications_request_service

import (
	"context"

	"github.com/biancarosa/wow-realm-status-notifier/models"
)

type NotificationReqServiceInterface interface {
	Create(*models.NotificationsRequest) error
}

type notificationsRequestService struct {
	ctx context.Context
}

func New(ctx context.Context) *notificationsRequestService {
	s := new(notificationsRequestService)
	s.ctx = ctx
	return s
}
