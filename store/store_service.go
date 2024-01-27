package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)


type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

// for cache lifetime. Should be based on LRU (least Recenly Used) policy, but for the sake of simplicity,
// we used time duration.
const CacheDuration = 6 * time.Hour


// initialize store service
func InitializeStore() *StorageService{
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0, 
	})

	// check redis connectivity
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init redis: %v", err))
	}

	fmt.Printf("\nRedis started: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// for the saving between original and generated short URL 
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

// reverse the short url to the longer, initial URL so that we can redirect it.
func RetrieveInitialUrl(shortUrl string) string {
	initialUrl, err := storeService.redisClient.Get(ctx, shortUrl).Result() 
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return initialUrl
}