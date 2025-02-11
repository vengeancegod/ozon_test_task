package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/internal/service"
)

type App struct {
	urlService service.URLService
}

func NewApp(urlService service.URLService) *App {
	return &App{urlService: urlService}
}

func (app *App) handleOriginalToShort(w http.ResponseWriter, r *http.Request) {
	var req models.ShortRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, models.ErrInvReq, http.StatusBadRequest)
		return
	}

	resp, err := app.urlService.OriginalToShort(req)
	if err != nil {
		http.Error(w, models.ErrGenShort, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (app *App) handleShortToOriginal(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Query().Get("short")
	if shortURL == "" {
		http.Error(w, models.ErrMissShortURL, http.StatusBadRequest)
		return
	}

	resp, err := app.urlService.ShortToOriginal(shortURL)
	if err != nil {
		http.Error(w, models.ErrShortNotFound, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ots", app.handleOriginalToShort)
	mux.HandleFunc("/sto", app.handleShortToOriginal)
	return mux
}
