package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"os"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	switch {
	case strings.HasPrefix(body.Message.Text, "/notify-status"): notifyStatus(body)
	case strings.HasPrefix(body.Message.Text, "/stop-notifying"): stopNotifying(body)
	default: return
	}
}

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func notifyStatus(body *webhookReqBody) error {
	return sendMessage(body.Message.Chat.ID, "Registered! Will notify you when server is up or down!")
}

func stopNotifying(body *webhookReqBody) error {
	return sendMessage(body.Message.Chat.ID, "I will stop notifying you when server is up or down!")
}

// sendMessage takes a chatID and sends a message to them
func sendMessage(chatID int64, message string) error {
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
	token := os.Getenv("TELEGRAM_TOKEN")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

// FInally, the main funtion starts our server on port 3000
func main() {
	fmt.Println("Running server")
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}