package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	mu       *sync.Mutex
	entries  map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
		mu:       &sync.Mutex{},
		entries:  make(map[string]cacheEntry),
	}
	cache.reapLoop()
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
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for {
			if _, ok := <-ticker.C; !ok {
				break
			}
			c.mu.Lock()
			now := time.Now()
			for key, entry := range c.entries {
				if now.Sub(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
