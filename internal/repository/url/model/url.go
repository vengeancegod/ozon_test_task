package url

type ShortRequest struct {
	OriginalURL string `json:"url"`
}

type ShortResponse struct {
	ShortURL string `json:"short_url"`
}