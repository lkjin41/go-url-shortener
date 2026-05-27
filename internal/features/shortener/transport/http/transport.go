package shortener_transport_http

import "github.com/lkjin41/go-url-shortener/internal/features/shortener/service"

// ShortenerHTTPHandler is the HTTP handler for the shortener feature.
type ShortenerHTTPHandler struct {
	shortenerService *service.ShortenerService
}

// NewShortenerHTTPHandler creates a new instance of ShortenerHTTPHandler.
func NewShortenerHTTPHandler(shortenerService *service.ShortenerService) *ShortenerHTTPHandler {
	return &ShortenerHTTPHandler{
		shortenerService: shortenerService,
	}
}
