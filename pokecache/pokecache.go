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
	mu    sync.Mutex
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
	cache.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	cache.data[key] = entry
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	val, ok := cache.data[key]
	if !ok {
		return nil, false
	}
	
	return val.val, true
}


func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)	
	for range ticker.C {
		cache.mu.Lock()
		for key, val  := range cache.data {
			if time.Since(val.createdAt) > cache.interval {
				delete(cache.data, key)
			}
		}
		cache.mu.Unlock()
	}
}