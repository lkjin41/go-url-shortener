package service

import "context"

// GetInitialUrl retrieves the original URL corresponding to the provided short link
func (s *ShortenerService) GetInitialUrl(ctx context.Context, shortLink string) (string, error) {
	var originalUrl string
	if err := s.repo.GetOriginalUrl(ctx, shortLink, &originalUrl); err != nil {
		return "", err
	}
	return originalUrl, nil
}
