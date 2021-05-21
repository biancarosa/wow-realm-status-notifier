package handlers

import (
	"fmt"
)

func StopNotifyingHandler(body *webhookReqBody) {
	fmt.Println("Notify status")
	// TODO: Remove data from collection
	sendMessage(body.Message.Chat.ID, "I will stop notifying you when server is up or down!")
}
