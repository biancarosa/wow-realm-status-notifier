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
	server := http.Server{
		Addr: fmt.Sprintf(":%s", config.Port),
	}

	database.Init("local.db")

	fmt.Println("Running server on port", config.Port)
	h := handlers.New()
	http.HandleFunc("/", h.DependenciesMiddleware(h.MainHandler))
	http.HandleFunc("/request-id", h.DependenciesMiddleware(h.RequestIDHandler))
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Shut down")
}
