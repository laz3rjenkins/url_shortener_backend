package shortener

type Shorten struct {
	Url string `form:"url" binding:"required"`
}

type ShortenUrlAttributes struct {
	OriginalURL   string `json:"original_url"`
	ShortenURL    string `json:"shorten_url"`
	RedirectCount int    `json:"redirect_count"`
}
