package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/biancarosa/wow-realm-status-notifier/configuration"
	"github.com/biancarosa/wow-realm-status-notifier/services"
)

var config *configuration.Config

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID uint8 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type Handler struct {
	Services *services.AppServices
}

// This handler is called everytime telegram sends us a webhook event
func MainHandler(w http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
	fmt.Println("received message", body.Message.Text)

	services := services.GetServices(req.Context())
	h := &Handler{
		Services: services,
	}
	switch {
	case strings.HasPrefix(body.Message.Text, "/notify-status"):
		h.NotifyStatusHandler(body)
	case strings.HasPrefix(body.Message.Text, "/stop-notifying"):
		h.StopNotifyingHandler(body)
	default:
		return
	}
}
