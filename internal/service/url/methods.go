package url

import (
	"url-shortener/internal/models"
	"url-shortener/pkg/generator"
)

func (s *Service) OriginalToShort(sRequest models.ShortRequest) (models.ShortResponse, error) {
	var shortURL string
	for {
		shortURL = generator.GenerateShortURL()
		if _, err := s.urlRepository.ShortToOriginal(shortURL); err != nil {
			break
		}
	}
	
	return s.urlRepository.OriginalToShort(sRequest, shortURL)
}

func (s *Service) ShortToOriginal(shortURL string) (models.ShortResponse, error) {
	return s.urlRepository.ShortToOriginal(shortURL)
}