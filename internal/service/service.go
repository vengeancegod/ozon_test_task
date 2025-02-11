package service

import "url-shortener/internal/models"

type URLService interface {
	OriginalToShort(sRequest models.ShortRequest) (models.ShortResponse, error)
	ShortToOriginal(shortURL string) (models.ShortResponse, error)
}
