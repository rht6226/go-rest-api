package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rht6226/go-rest-api/entity"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

// Constructor for a new redis cache
func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

// create a new Redis client
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

// set a key in redis
func (cache *redisCache) Set(key string, value *entity.Post) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

// get key
func (cache *redisCache) Get(key string) *entity.Post {
	client := cache.getClient()

	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post
}
