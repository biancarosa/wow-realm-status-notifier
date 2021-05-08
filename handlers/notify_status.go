package handlers

import (
	"fmt"
)

func NotifyStatusHandler(body *webhookReqBody) {
	fmt.Println("Notify status")
	sendMessage(body.Message.Chat.ID, "Registered! Will notify you when server is up or down!")
}
