package url 

import (
	"url-shortener/internal/repository"
)

type Service struct {
	urlRepository repository.URLRepository
}

func NewService(urlRepository repository.URLRepository) *Service {
	return &Service{
		urlRepository: urlRepository,
	}
}