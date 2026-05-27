package service

// ShortenerService provides the business logic for URL shortening operations
type ShortenerService struct {
	repo ShortenerRepository
}

// ShortenerRepository defines the interface for the repository layer
type ShortenerRepository interface {
	SaveUrlMapping(key string, value any) error
	GetOriginalUrl(key string) (string, error)
}

// NewShortenerService creates a new instance of ShortenerService
func NewShortenerService(repo ShortenerRepository) *ShortenerService {
	return &ShortenerService{
		repo: repo,
	}
}
