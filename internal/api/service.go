package api

import (
	"url-shortener/internal/service"
	desc "url-shortener/pkg/urlshortener_v3"
)

type Implementation struct {
	desc.UnimplementedURLShortenerServer
	URLService service.URLService
}

func NewImplementation(urlService service.URLService) *Implementation {
	return &Implementation{
		URLService: urlService,
	}
}
