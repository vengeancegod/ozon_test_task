package url

import (
	"database/sql"
	"url-shortener/internal/models"
	rep "url-shortener/internal/repository"
	_ "github.com/lib/pq"
)

var _ rep.URLRepository = (*repositoryPG)(nil)

type repositoryPG struct {
	DB *sql.DB
}

func NewRepositoryPG(db *sql.DB) (*repositoryPG, error) {
	return &repositoryPG{
		DB: db,
	}, nil
}

func (repo *repositoryPG) OriginalToShort(sRequest models.ShortRequest, shortURL string) (models.ShortResponse, error) {
	_, err := repo.DB.Exec("INSERT INTO urls (original_url, short_url) VALUES ($1, $2)", sRequest.OriginalURL, shortURL)
	if err != nil {
		return models.ShortResponse{}, err
	}
	return models.ShortResponse{ShortURL: shortURL}, nil
}

func (repo *repositoryPG) ShortToOriginal(shortURL string) (models.ShortRequest, error) {
	var originalURL string
	err := repo.DB.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows{
			return models.ShortRequest{}, err
		}
		return models.ShortRequest{}, err
	}
	return models.ShortRequest{OriginalURL: originalURL}, nil
}

