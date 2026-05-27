package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) SaveUrlMapping(ctx context.Context, key string, value any) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockRepo) GetOriginalUrl(ctx context.Context, key string, dest any) error {
	args := m.Called(ctx, key, dest)
	return args.Error(0)
}

const UserID = "134042d4-96da-46da-95b3-36f939f621d4"

func TestCreateShortLink_SavesMappingAndReturnsShortLink(t *testing.T) {
	repo := new(MockRepo)
	service := NewShortenerService(repo)

	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	expectedShort := generateShortLink(initialLink, UserID)

	repo.On("SaveUrlMapping", mock.Anything, expectedShort, initialLink).Return(nil)

	shortLink, err := service.CreateShortLink(context.Background(), initialLink, UserID)

	assert.NoError(t, err)
	assert.Equal(t, expectedShort, shortLink)
	repo.AssertExpectations(t)
}

func TestCreateShortLink_ReturnsErrorWhenRepoFails(t *testing.T) {
	repo := new(MockRepo)
	service := NewShortenerService(repo)

	initialLink := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	saveErr := errors.New("save failed")

	repo.On("SaveUrlMapping", mock.Anything, mock.Anything, initialLink).Return(saveErr)

	shortLink, err := service.CreateShortLink(context.Background(), initialLink, UserID)

	assert.Error(t, err)
	assert.Empty(t, shortLink)
	repo.AssertExpectations(t)
}
