package service

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// CreateShortLink generates a short link for the given initial link and user ID,
// and saves the mapping in the repository
func (s *ShortenerService) CreateShortLink(initialLink string, userId string) (string, error) {
	shortLink := generateShortLink(initialLink, userId)
	err := s.repo.SaveUrlMapping(shortLink, initialLink)
	if err != nil {
		return "", fmt.Errorf("failed to save URL mapping: %w", err)
	}
	return shortLink, nil
}

func generateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
