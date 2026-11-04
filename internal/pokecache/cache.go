package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]CacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)

	newCache := &Cache{
		entries:  map[string]CacheEntry{},
		interval: interval,
		mu:       sync.Mutex{},
	}

	go func() {
		for range ticker.C {
			newCache.reaploop(<-ticker.C)
		}
	}()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	ok := c.mu.TryLock()
	if !ok {
		return nil, false
	}
	defer c.mu.Unlock()

	if c.entries[key].val != nil {
		return c.entries[key].val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reaploop(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, entry := range c.entries {
		if t.Sub(entry.createdAt) >= c.interval {
			c.entries[k] = CacheEntry{}
		}
	}
}
