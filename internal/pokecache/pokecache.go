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
	mu       sync.Mutex
	interval time.Duration
	entries  map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	cutoff := time.Now().Add(-c.interval)

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}
