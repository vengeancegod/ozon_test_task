package main

import (
	"flag"
	"log"
	"net"
	"os"
	"url-shortener/internal/api"
	"url-shortener/internal/infrastructure"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/url"
	uService "url-shortener/internal/service/url"
	desc "url-shortener/pkg/urlshortener_v3"

	"google.golang.org/grpc"
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
			log.Fatalf("Failed to initialize DB: %v", err)
		}
		repo, err = url.NewRepositoryPG(db)
		if err != nil {
			log.Fatalf("Failed to create PG repository: %v", err)
		}
	case "inmemory":
		repo = url.NewRepository()
	default:
		log.Fatalf("Unknown storage type: %s", *storageType)
	}

	urlService := uService.NewService(repo)
	// urlService := uService.NewImplementation(uService)

	server := grpc.NewServer()

	desc.RegisterURLShortenerServer(server, &api.Implementation{
		URLService: urlService,
	})

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server started on port :50051")

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
