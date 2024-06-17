package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data     map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache {
		data: make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}


func (cache *Cache) Add(key string, val []byte) {
	
}