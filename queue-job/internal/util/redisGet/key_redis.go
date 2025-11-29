package redisGet

import (
	"context"
	"queue-job-service/internal/config"

	"github.com/redis/go-redis/v9"
)


var Rdb *redis.Client


func GetByKey(key string) (string, error) {
	ctx := context.Background()
	fullKey := "test:key:" + key

	val, err := config.RedisClient.Get(ctx, fullKey).Result()
	if err != nil {
		return "", err 
	}

	return val, nil 
}