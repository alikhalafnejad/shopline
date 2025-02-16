package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"shopline/config"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	settings := config.LoadSettings()

	client := redis.NewClient(&redis.Options{
		Addr:     settings.RedisAddr,
		Password: settings.RedisPassword,
		DB:       settings.RedisDB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &RedisClient{Client: client}
}

func NewRedisClientWithPool(addr, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		PoolSize:   100, // Maximum number of connections in the pool
		MaxRetries: 3,   // Retry attempts for failed operations
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &RedisClient{Client: client}
}
