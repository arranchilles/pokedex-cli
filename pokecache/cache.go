package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]CacheEntry
	mutex   sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	Val       []byte
}

func (c *Cache) Add(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Entries[key] = CacheEntry{time.Now(), data}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if entry, ok := c.Entries[key]; ok {
		return entry.Val, true
	}
	return nil, false
}

func (c *Cache) reaploop(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for entry := range c.Entries {
		if time.Since(c.Entries[entry].createdAt) > interval {
			delete(c.Entries, entry)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	cache := &Cache{Entries: map[string]CacheEntry{}}

	go func() {
		for range ticker.C {
			cache.reaploop(interval)
		}
	}()
	return cache
}
