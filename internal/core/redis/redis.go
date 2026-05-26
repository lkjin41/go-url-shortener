package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	coreconfig "github.com/lkjin41/go-url-shortener/internal/core/config"
	"github.com/redis/go-redis/v9"
)

const DefaultTTL = 15 * time.Minute

// StorageService define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// InitializeStore initializing the store service and return a store pointer
func InitializeStore(ctx context.Context, cfg *coreconfig.RedisConfig) *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	return &StorageService{redisClient: redisClient}
}

// Set value in Redis with default TTL
func (s *StorageService) Set(ctx context.Context, key string, value any) error {

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value for key %s: %w", key, err)
	}

	err = s.redisClient.Set(ctx, key, data, DefaultTTL).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s in Redis: %w", key, err)
	}

	return nil
}

// Get retrieve value from Redis
func (s *StorageService) Get(ctx context.Context, key string, dest any) error {

	value, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return fmt.Errorf("key %s not found in Redis: %w", key, err)
		}
		return fmt.Errorf("failed to get key %s from Redis: %w", key, err)
	}

	err = json.Unmarshal([]byte(value), dest)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value for key %s: %w", key, err)
	}

	return nil
}
