package shortener_redis_repository

import (
	"context"
	"errors"

	coreerrors "github.com/lkjin41/go-url-shortener/internal/core/errors"
	"github.com/redis/go-redis/v9"
)

// Repository defines the interface for interacting with the Redis storage service for URL mappings
type Repository struct {
	service StorageEngine
}

// StorageEngine defines the interface for the storage engine used by the Repository
type StorageEngine interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string, dest any) error
}

// NewRepository creates a new instance of Repository with the provided storage service
func NewRepository(service StorageEngine) *Repository {
	return &Repository{service: service}
}

// SaveUrlMapping saves the mapping between the original URL and the shortened URL in Redis
func (r *Repository) SaveUrlMapping(ctx context.Context, key string, value any) error {
	err := r.service.Set(ctx, key, value)
	if err != nil {
		return coreerrors.NewInternalServerError("failed to save URL mapping in Redis", err)
	}
	return nil
}

// GetOriginalUrl retrieves the original URL based on the shortened URL from Redis
func (r *Repository) GetOriginalUrl(ctx context.Context, key string, dest any) error {
	err := r.service.Get(ctx, key, dest)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return coreerrors.NewNotFoundError("short link does not exist", err)
		}

		return coreerrors.NewInternalServerError("failed to get url from storage", err)
	}

	return nil
}
