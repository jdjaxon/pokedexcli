package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries      map[string]cacheEntry
	mu           sync.RWMutex
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:      map[string]cacheEntry{},
		mu:           sync.RWMutex{},
		reapInterval: interval,
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, entry []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       entry,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > c.reapInterval {
			delete(c.entries, key)
		}
	}
}
