package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry    map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entry:    make(map[string]cacheEntry),
		mu:       sync.RWMutex{},
		interval: interval,
	}
	go (&cache).reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{time.Now(), val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	v, ok := c.entry[key]
	c.mu.RUnlock()
	if ok {
		return v.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k := range c.entry {
			diff := time.Since(c.entry[k].createdAt)
			if diff > c.interval {
				delete(c.entry, k)
			}
		}
		c.mu.Unlock()
	}
}
