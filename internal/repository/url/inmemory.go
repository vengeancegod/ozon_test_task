package url

import (
	"errors"
	"sync"
	"url-shortener/internal/models"
	rep "url-shortener/internal/repository"
	repoModel "url-shortener/internal/repository/url/model"
)

var _ rep.URLRepository = (*repository)(nil)

type repository struct {
	sRequest  map[string]*repoModel.ShortRequest
	sResponse map[string]*repoModel.ShortResponse
	mu        sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		sRequest:  make(map[string]*repoModel.ShortRequest),
		sResponse: make(map[string]*repoModel.ShortResponse),
	}
}

func (repo *repository) OriginalToShort(sRequest models.ShortRequest, shortURL string) (models.ShortResponse, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for short, req := range repo.sRequest {
		if req.OriginalURL == sRequest.OriginalURL {
			if response, exists := repo.sResponse[short]; exists {
				return models.ShortResponse{ShortURL: response.ShortURL}, nil
			}
		}
	}

	repo.sRequest[shortURL] = &repoModel.ShortRequest{OriginalURL: sRequest.OriginalURL}
	repo.sResponse[shortURL] = &repoModel.ShortResponse{ShortURL: shortURL}

	return models.ShortResponse{ShortURL: shortURL}, nil
}

func (repo *repository) ShortToOriginal(shortURL string) (models.ShortRequest, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if req, exists := repo.sRequest[shortURL]; exists {
		return models.ShortRequest{OriginalURL: req.OriginalURL}, nil
	}

	return models.ShortRequest{}, errors.New(models.ErrNotFound)
}
