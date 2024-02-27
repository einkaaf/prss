package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var once sync.Once
var pCacheInstance *PrssCache

type PrssCache struct {
	c *cache.Cache
}

func NewPrssCache() *PrssCache {

	once.Do(func() {
		cache := cache.New(5*time.Minute, 10*time.Minute)
		pCacheInstance = &PrssCache{c: cache}
	})
	return pCacheInstance
}

func (pc *PrssCache) SetCachedItem(key string, item any) {
	pc.c.Set(key, item, time.Minute*5)
}

func (pc *PrssCache) GetCachedItem(key string) (interface{}, bool) {
	return pc.c.Get(key)
}
