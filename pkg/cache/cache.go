package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

func NewCache(ttl time.Duration) *Cache {
	redisAddress := os.Getenv("REDIS_ADDR")

	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		err := fmt.Errorf("could not create cache: %w", err)
		fmt.Printf("[ERROR] %s\n", err.Error())
	}

	return &Cache{client: client, ttl: ttl}
}

func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
