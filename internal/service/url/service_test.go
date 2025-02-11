package url

import (
  "testing"
  "url-shortener/internal/models"
	"url-shortener/internal/repository/url"
  "github.com/stretchr/testify/assert"
)

func TestOriginalToShort(t *testing.T) {
  repo := url.NewRepository()
  service := NewService(repo)

  req := models.ShortRequest{OriginalURL: "https://pepepepe.ru"}

  resp, err := service.OriginalToShort(req)

  assert.NoError(t, err)
  assert.NotEmpty(t, resp.ShortURL) 
}

func TestShortToOriginal(t *testing.T) {
  repo := url.NewRepository()
  service := NewService(repo)

  req := models.ShortRequest{OriginalURL: "https://qBuEIc1"}
  shortURL := "abc123"
  repo.OriginalToShort(req, shortURL)

  resp, err := service.ShortToOriginal(shortURL)

  assert.NoError(t, err)
  assert.Equal(t, "https://qwere.com", resp.ShortURL)
}