package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	config "github.com/lkjin41/go-url-shortener/internal/core/config"
	core_redis "github.com/lkjin41/go-url-shortener/internal/core/redis"
	shortener_redis_repository "github.com/lkjin41/go-url-shortener/internal/features/shortener/repository/redis"
	shortener_service "github.com/lkjin41/go-url-shortener/internal/features/shortener/service"
	shortener_transport_http "github.com/lkjin41/go-url-shortener/internal/features/shortener/transport/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	cfg := config.Load()

	redisClient := core_redis.InitializeStore(context.Background(), &cfg.Redis)
	redisRepo := shortener_redis_repository.NewRepository(redisClient)
	service := shortener_service.NewShortenerService(redisRepo)
	handler := shortener_transport_http.NewShortenerHTTPHandler(service)

	r.POST("/create-short-url", handler.CreateShortURL)

	r.GET("/:shortUrl", handler.HandleShortUrlRedirect)

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
