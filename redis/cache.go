package redis

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	MemCache *cache.Cache
)

func init()  {
	MemCache = cache.New(time.Hour * 1, time.Hour * 2)
}

func hasCacheKey(key string) bool  {
	_,ok := MemCache.Get(key)
	if !ok {
		MemCache.Set(key, 1, cache.DefaultExpiration)
		return false
	}
	return true
}