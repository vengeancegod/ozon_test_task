package main

import (
	"log"
	"net/http"
	"url-shortener/internal/handlers"
	"url-shortener/internal/repository/url"
	uService "url-shortener/internal/service/url"
)

func main() {
	urlRepo := url.NewRepository()
	urlService := uService.NewService(urlRepo)
	app := handlers.NewApp(urlService)

	server := http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}

	log.Println(("Listen on :8080"))
	server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
