package url

type ShortRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortResponse struct {
	ShortURL string `json:"short_url"`
}
