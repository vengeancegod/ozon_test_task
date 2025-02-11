package url

import (
	"testing"
	"url-shortener/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestOriginalToShort(t *testing.T) {
	repo := NewRepository()
	req := models.ShortRequest{OriginalURL: "https://vk.com/didididid"}
	shortURL := "abc123"

	resp, err := repo.OriginalToShort(req, shortURL)

	assert.NoError(t, err)
	assert.Equal(t, shortURL, resp.ShortURL)
}

func TestShortToOriginal(t *testing.T) {
	repo := NewRepository()
	req := models.ShortRequest{OriginalURL: "https://vk.com/didididid"}
	shortURL := "abc123"
	repo.OriginalToShort(req, shortURL)

	resp, err := repo.ShortToOriginal(shortURL)

	assert.NoError(t, err)
	assert.Equal(t, "https://vk.com/didididid", resp.ShortURL)
}
