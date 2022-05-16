package cache

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const (
	CacheDuration = 6 * time.Hour
)

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\nRedis started: pong message = {%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shorturl string, originalurl string, userId string) {
	err := storeService.redisClient.Set(ctx, shorturl, originalurl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shorturl, originalurl))
	}
}

func RetrieveInitialUrl(shorturl string) string {
	result, err := storeService.redisClient.Get(ctx, shorturl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shorturl))
	}
	return result
}
