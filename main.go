package main

import (
	"MidhunRajeevan/payments/app/routes"
	"MidhunRajeevan/payments/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitializeApp()
	config.InitializeDB()

	url := fmt.Sprintf(":%d", config.App.ListenPort)
	log.Println("Starting server at " + url)

	srv := &http.Server{
		Addr:    url,
		Handler: routes.Routes(),
	}
	log.Fatal(srv.ListenAndServe())

}
