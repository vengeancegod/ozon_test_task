package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/models"
	uRepo "url-shortener/internal/repository/url"
	uService "url-shortener/internal/service/url"

	"github.com/stretchr/testify/assert"
)

func TestHandleOriginalToShort(t *testing.T) {

	fakeRepo := uRepo.NewRepository()
	fakeService := uService.NewService(fakeRepo)
	app := NewApp(fakeService)

	reqBody, _ := json.Marshal(models.ShortRequest{OriginalURL: "https://example.com"})
	req := httptest.NewRequest(http.MethodPost, "/ots", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.handleOriginalToShort(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var resp models.ShortResponse
	json.NewDecoder(rec.Body).Decode(&resp)
	assert.NotEmpty(t, resp.ShortURL) 
}

func TestHandleShortToOriginal(t *testing.T) {

	fakeRepo := uRepo.NewRepository()
	fakeService := uService.NewService(fakeRepo)
	app := NewApp(fakeService)

	originalURL := "https://vk.com/isdfgsjdfkgjsfgjjgfs"
	shortURL := "abc123"
	fakeRepo.OriginalToShort(models.ShortRequest{OriginalURL: originalURL}, shortURL)

	req := httptest.NewRequest(http.MethodGet, "/sto?short="+shortURL, nil)
	rec := httptest.NewRecorder()

	app.handleShortToOriginal(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var resp models.ShortResponse
	json.NewDecoder(rec.Body).Decode(&resp)
	assert.Equal(t, originalURL, resp.ShortURL)
}
