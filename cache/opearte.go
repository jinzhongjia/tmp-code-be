package cache

import (
	"context"
	"time"

	"github.com/eko/gocache/v2/store"
)

var ctx = context.Background()

func (that *Cache) Set(key interface{}, value interface{}, expire time.Duration) error {
	return that.C.Set(ctx, key, value, &store.Options{
		Expiration: expire,
	})
}

func (that *Cache) Get(key interface{}) (interface{}, error) {
	return that.C.Get(ctx, key)
}

func (that *Cache) Del(key interface{}) error {
	return that.C.Delete(ctx, key)
}
