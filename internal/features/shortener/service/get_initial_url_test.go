package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInitialUrl_ReturnsOriginal(t *testing.T) {
	repo := new(MockRepo)
	service := NewShortenerService(repo)

	shortLink := "jTa4L57P"
	original := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"

	repo.On("GetOriginalUrl", shortLink).Return(original, nil)

	result, err := service.GetInitialUrl(shortLink)

	assert.NoError(t, err)
	assert.Equal(t, original, result)
	repo.AssertExpectations(t)
}

func TestGetInitialUrl_PropagatesError(t *testing.T) {
	repo := new(MockRepo)
	service := NewShortenerService(repo)

	shortLink := "d66yfx7N"
	fetchErr := errors.New("not found")

	repo.On("GetOriginalUrl", shortLink).Return("", fetchErr)

	result, err := service.GetInitialUrl(shortLink)

	assert.Error(t, err)
	assert.Empty(t, result)
	repo.AssertExpectations(t)
}
