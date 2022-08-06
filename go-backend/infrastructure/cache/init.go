package cache

import (
	"go-backend/config"
	"go-backend/infrastructure/cache/redis"
	"go-backend/interfaces/ports/cache"
	"log"
	"os"
)

// Constants
const (
	REDIS_CACHE string = "redis"
)

var cacheType string

// Cache
var cacheObj cache.ICache

func InitializeCache() {
	cacheType = os.Getenv(config.CACHE_TYPE)
	if cacheType == REDIS_CACHE {
		redis.GetCache()
		log.Printf("%s cache initialized successfully", cacheType)
	} else {
		log.Fatalln("CACHE_TYPE could not match.")
	}
}

func GetCache() cache.ICache {
	if cacheObj == nil {
		if cacheType == REDIS_CACHE {
			cacheObj = redis.NewCache()
		}
	}
	return cacheObj
}
