package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"url-shortener/internal/handlers"
	"url-shortener/internal/infrastructure"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/url"
	uService "url-shortener/internal/service/url"
)

func main() {
	storageType := flag.String("storage", "inmemory", "Storage type to using")
	flag.Parse()

	var repo repository.URLRepository
	if envStorageType := os.Getenv("STORAGE+TYPE"); envStorageType != "" {
		*storageType = envStorageType
	}
	switch *storageType {
	case "postgres":
		db, err := infrastructure.InitDB()
		if err != nil {
			log.Fatal("Failed initialize db: %v", err)
		}
		repo, err = url.NewRepositoryPG(db)
		if err != nil {
			log.Fatalf("Failed create PG repository: %v", err)
		}
	case "inmemory":
		repo = url.NewRepository()
	default:
		log.Fatalf("Unknown storage type: %s", *storageType)
	}

	// urlRepo := url.NewRepository()
	urlService := uService.NewService(repo)
	app := handlers.NewApp(urlService)

	server := http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}

	log.Println(("Listen on :8080"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
