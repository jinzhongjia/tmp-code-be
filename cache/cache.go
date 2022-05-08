package cache

import (
	"log"
	"time"
	"tmp-code/flag"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gocache "github.com/patrickmn/go-cache"
)

type Cache struct {
	C *cache.Cache
}

func NewCache() *Cache {
	switch flag.T {
	case "redis":
		return &Cache{
			C: newRedis(),
		}
	case "memory":
		return &Cache{
			C: newBigCache(),
		}
	}
	log.Fatal("the type of cache is error!")
	return &Cache{}
}

func newBigCache() *cache.Cache {
	gocacheClient := gocache.New(time.Hour, 10*time.Minute)
	gocacheStore := store.NewGoCache(gocacheClient, nil) // No options provided (as second argument)
	return cache.New(gocacheStore)
}

func newRedis() *cache.Cache {
	redisStore := store.NewRedis(redis.NewClient(&redis.Options{
		Addr:     flag.RedisAddr,
		Username: flag.RedisUserName,
		Password: flag.RedisPasswd,
		// 默认redis有16个数据库，这里使用15
		DB: 15,
	}), nil)

	return cache.New(redisStore)
}
