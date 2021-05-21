package main

import (
	"fmt"
	"net/http"

	"github.com/biancarosa/wow-realm-status-notifier/configuration"
	"github.com/biancarosa/wow-realm-status-notifier/database"
	"github.com/biancarosa/wow-realm-status-notifier/handlers"
)

var config *configuration.Config

func main() {
	config = configuration.GetConfig()

	database.Init("local.db")

	fmt.Println("Running server on port", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), http.HandlerFunc(handlers.MainHandler))
}
