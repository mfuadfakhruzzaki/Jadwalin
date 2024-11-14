package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// InitRedis initializes the Redis connection
func InitRedis() error {
	redisHost := GetEnv("REDIS_HOST", "localhost")
	redisPort := GetEnv("REDIS_PORT", "6379")
	redisPassword := GetEnv("REDIS_PASSWORD", "")

	// Create Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,  // No password if empty
		DB:       0,              // Use default DB
	})

	// Ping Redis to verify the connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Error connecting to Redis: %v", err)
		return err
	}

	log.Println("Successfully connected to Redis")
	return nil
}

// CloseRedis properly closes the Redis client connection
func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	} else {
		log.Println("Redis connection closed")
	}
}
