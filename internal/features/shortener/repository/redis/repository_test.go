package shortener_redis_repository

import (
	"context"
	"errors"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	coreerrors "github.com/lkjin41/go-url-shortener/internal/core/errors"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Set(ctx context.Context, key string, value any) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockStorage) Get(ctx context.Context, key string, dest any) error {
	args := m.Called(ctx, key, dest)
	return args.Error(0)
}

func assertAppError(t *testing.T, err error, statusCode int, message string, wrapped error) {
	t.Helper()

	var appErr *coreerrors.AppError
	if !assert.Error(t, err) {
		return
	}
	if !assert.True(t, errors.As(err, &appErr)) {
		return
	}

	assert.Equal(t, statusCode, appErr.StatusCode)
	assert.Equal(t, message, appErr.Message)
	if wrapped != nil {
		assert.True(t, errors.Is(err, wrapped))
	}
}

func TestRepository_SaveUrlMapping(t *testing.T) {
	ctx := context.Background()
	errBoom := errors.New("boom")

	tests := []struct {
		name           string
		setup          func(ms *MockStorage)
		expectedErrMsg string
		expectedCode   int
		wrappedErr     error
	}{
		{
			name: "Success",
			setup: func(ms *MockStorage) {
				ms.On("Set", mock.Anything, "short:abc", "http://example.com").Return(nil)
			},
			expectedErrMsg: "",
		},
		{
			name: "Storage failure",
			setup: func(ms *MockStorage) {
				ms.On("Set", mock.Anything, "short:abc", "http://example.com").Return(errBoom)
			},
			expectedErrMsg: "failed to save URL mapping in Redis",
			expectedCode:   500,
			wrappedErr:     errBoom,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := new(MockStorage)
			tt.setup(ms)

			repo := NewRepository(ms)
			err := repo.SaveUrlMapping(ctx, "short:abc", "http://example.com")

			if tt.expectedErrMsg == "" {
				assert.NoError(t, err)
			} else {
				assertAppError(t, err, tt.expectedCode, tt.expectedErrMsg, tt.wrappedErr)
			}

			ms.AssertExpectations(t)
		})
	}
}

func TestRepository_GetOriginalUrl(t *testing.T) {
	ctx := context.Background()
	errBoom := errors.New("boom")

	tests := []struct {
		name           string
		key            string
		setup          func(ms *MockStorage)
		expectedErrMsg string
		expectedCode   int
		wrappedErr     error
		expectedURL    string
	}{
		{
			name: "Success",
			key:  "short:ok",
			setup: func(ms *MockStorage) {
				ms.On("Get", mock.Anything, "short:ok", mock.Anything).
					Run(func(args mock.Arguments) {
						dest := args.Get(2).(*string)
						*dest = "http://example.com"
					}).
					Return(nil)
			},
			expectedURL: "http://example.com",
		},
		{
			name: "Not found",
			key:  "short:missing",
			setup: func(ms *MockStorage) {
				ms.On("Get", mock.Anything, "short:missing", mock.Anything).Return(redis.Nil)
			},
			expectedErrMsg: "short link does not exist",
			expectedCode:   404,
			wrappedErr:     redis.Nil,
		},
		{
			name: "Storage failure",
			key:  "short:err",
			setup: func(ms *MockStorage) {
				ms.On("Get", mock.Anything, "short:err", mock.Anything).Return(errBoom)
			},
			expectedErrMsg: "failed to get url from storage",
			expectedCode:   500,
			wrappedErr:     errBoom,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := new(MockStorage)
			tt.setup(ms)

			repo := NewRepository(ms)

			var dest string
			err := repo.GetOriginalUrl(ctx, tt.key, &dest)

			if tt.expectedErrMsg == "" {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedURL, dest)
			} else {
				assertAppError(t, err, tt.expectedCode, tt.expectedErrMsg, tt.wrappedErr)
			}

			ms.AssertExpectations(t)
		})
	}
}
