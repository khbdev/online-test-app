package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	testpb "github.com/khbdev/proto-online-test/proto/test"
	"github.com/redis/go-redis/v9"
)

type TestRepository struct {
	RedisClient *redis.Client
	Ctx         context.Context
}


type TestData struct {
    Name     string            `json:"name"`
    Sections []*testpb.Section `json:"sections"` 
}


func NewRepository(rdb *redis.Client) *TestRepository {
	return &TestRepository{
		RedisClient: rdb,
		Ctx:         context.Background(),
	}
}


func (r *TestRepository) Set(name, key string, data TestData, ttl time.Duration) error {

	if err := r.RedisClient.Set(r.Ctx, "test:name:"+name, key, ttl).Err(); err != nil {
		return fmt.Errorf("failed to set name->key: %w", err)
	}

	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal test data: %w", err)
	}

	if err := r.RedisClient.Set(r.Ctx, "test:key:"+key, jsonData, ttl).Err(); err != nil {
		return fmt.Errorf("failed to set key->value: %w", err)
	}

	return nil
}


func (r *TestRepository) Get(key string) (*TestData, error) {
	val, err := r.RedisClient.Get(r.Ctx, "test:key:"+key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get key: %w", err)
	}

	var data TestData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &data, nil
}
