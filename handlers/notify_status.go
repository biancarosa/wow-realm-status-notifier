package handlers

import (
	"fmt"

	"github.com/biancarosa/wow-realm-status-notifier/models/notifications_request"
	"github.com/biancarosa/wow-realm-status-notifier/services/notifications_request_service"
)

func NotifyStatusHandler(body *webhookReqBody) {
	fmt.Println("Notify status")
	s := notifications_request_service.New()
	nr := notifications_request.New("", "")
	// TODO: Add logging lib and log and handle this error properly
	err := s.Create(nr)
	if err != nil {
		sendMessage(body.Message.Chat.ID, "Something weird happenned in my circuits. Try again later, pretty please?")
	} else {
		sendMessage(body.Message.Chat.ID, "Registered! Will notify you when server is up or down!")
	}
}
