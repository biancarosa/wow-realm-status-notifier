package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/biancarosa/wow-realm-status-notifier/configuration"
)

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID uint8  `json:"chat_id"`
	Text   string `json:"text"`
}

// sendMessage takes a chatID and sends a message to them
func sendMessage(chatID uint8, message string) error {
	config = configuration.GetConfig()

	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.TelegramToken)
	fmt.Println("Sending message to telegram chat id", chatID)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
