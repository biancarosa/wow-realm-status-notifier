package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/biancarosa/wow-realm-status-notifier/models/notifications_request"
	"github.com/biancarosa/wow-realm-status-notifier/services"
)

func getServerFromText(text string) (string, error) {
	spplited := strings.Split(text, " ")
	if len(spplited) < 2 {
		fmt.Println(text)
		return "", errors.New("bad input")
	}
	return spplited[1], nil
}

func NotifyStatusHandler(body *webhookReqBody) error {
	fmt.Println("Notify status")
	serviceContainer := services.GetServices()
	fmt.Printf("%#v\n", body)
	server, err := getServerFromText(body.Message.Text)
	if err != nil {
		return sendMessage(body.Message.Chat.ID, "Something's wrong with your input.")
	}
	nr := notifications_request.New(body.Message.Chat.ID, server)
	// TODO: Add logging lib and log and handle this error properly
	err = serviceContainer.NotificationsRequestService.Create(nr)
	if err != nil {
		return sendMessage(body.Message.Chat.ID, "Something weird happenned in my circuits. Try again later, pretty please?")
	} else {
		return sendMessage(body.Message.Chat.ID, "Registered! Will notify you when server is up or down!")
	}
}
