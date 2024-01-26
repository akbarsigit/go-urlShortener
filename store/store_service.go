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

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}