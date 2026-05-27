package service

// GetInitialUrl retrieves the original URL corresponding to the provided short link
func (s *ShortenerService) GetInitialUrl(shortLink string) (string, error) {
	originalUrl, err := s.repo.GetOriginalUrl(shortLink)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}
