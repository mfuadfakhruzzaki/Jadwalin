package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
    redisHost := GetEnv("REDIS_HOST", "localhost")
    redisPort := GetEnv("REDIS_PORT", "6379")
    redisPassword := GetEnv("REDIS_PASSWORD", "")

    RedisClient = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
        Password: redisPassword,  // kosong jika tidak ada password
        DB:       0,              // gunakan database default
    })

    // Cek koneksi ke Redis
    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatal("Redis connection failed:", err)
    }

    log.Println("Redis connected successfully")
}
