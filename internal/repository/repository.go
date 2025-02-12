package repository

import "url-shortener/internal/models"

type URLRepository interface {
	OriginalToShort(sRequest models.ShortRequest, shortURL string) (models.ShortResponse, error)
	ShortToOriginal(shortURL string) (models.ShortRequest, error)
}