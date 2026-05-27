package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetInitialUrl_ReturnsOriginal(t *testing.T) {
	repo := new(MockRepo)
	service := NewShortenerService(repo)

	shortLink := "jTa4L57P"
	original := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"

	repo.On("GetOriginalUrl", mock.Anything, shortLink, mock.Anything).
		Run(func(args mock.Arguments) {
			dest := args.Get(2).(*string)
			*dest = original
		}).
		Return(nil)

	result, err := service.GetInitialUrl(context.Background(), shortLink)

	assert.NoError(t, err)
	assert.Equal(t, original, result)
	repo.AssertExpectations(t)
}

func TestGetInitialUrl_PropagatesError(t *testing.T) {
	repo := new(MockRepo)
	service := NewShortenerService(repo)

	shortLink := "d66yfx7N"
	fetchErr := errors.New("not found")

	repo.On("GetOriginalUrl", mock.Anything, shortLink, mock.Anything).Return(fetchErr)

	result, err := service.GetInitialUrl(context.Background(), shortLink)

	assert.Error(t, err)
	assert.Empty(t, result)
	repo.AssertExpectations(t)
}
