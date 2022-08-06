package redis

import (
	"context"
	"go-backend/config"
	"go-backend/interfaces/ports/cache"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

type Cache struct {
	rdb *redis.Client
}

func NewCache() cache.ICache {
	if rdb == nil {
		rdb = GetCache()
	}
	return &Cache{rdb: rdb}
}

func GetCache() *redis.Client {
	if rdb == nil {
		addr := os.Getenv(config.REDIS_ADDR)
		password := os.Getenv(config.REDIS_PASSWORD)
		db, err := strconv.Atoi(os.Getenv(config.REDIS_DB))
		if err != nil {
			log.Fatalf("Error parsing REDIS_DB parameter from .env file: %v\n", err)
		}
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
		log.Println("Redis cache created successfully")
	}
	return rdb
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expireSeconds int64) error {
	return c.rdb.Set(ctx, key, value, time.Duration(expireSeconds)*time.Second).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

func (c *Cache) Exists(ctx context.Context, key string) (bool, error) {
	val, err := c.rdb.Exists(ctx, key).Result()
	return val == 1, err
}
