package services

import "github.com/biancarosa/wow-realm-status-notifier/services/notifications_request_service"

type AppServices struct {
	NotificationsRequestService notifications_request_service.NotificationReqServiceInterface
}

var as *AppServices

func GetServices() *AppServices {
	if as == nil {
		as = &AppServices{
			NotificationsRequestService: notifications_request_service.New(),
		}
	}
	return as
}
