package cache

import (
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
)

type Cache struct {
	C *cache.Cache
}

func NewCache() *Cache {
	return &Cache{
		C: newBigCache(),
	}
}

func newBigCache() *cache.Cache {
	bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient, nil) // No options provided (as second argument)
	return cache.New(bigcacheStore)
}
