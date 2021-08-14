package main

import (
	"fmt"
	"net/http"

	"github.com/biancarosa/wow-realm-status-notifier/configuration"
	"github.com/biancarosa/wow-realm-status-notifier/database"
	"github.com/biancarosa/wow-realm-status-notifier/handlers"
)

func main() {
	config := configuration.GetConfig()
	server := http.Server{
		Addr: fmt.Sprintf(":%s", config.Port),
	}

	database.Init()

	fmt.Println("Running server on port", config.Port)
	http.HandleFunc("/", handlers.MainHandler)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Shutdown")
}
