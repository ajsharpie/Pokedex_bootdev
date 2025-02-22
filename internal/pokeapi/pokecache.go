package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval int) *Cache {
	c := &Cache{
		Entries:  make(map[string]cacheEntry),
		interval: time.Duration(interval) * time.Second,
		mu:       sync.Mutex{},
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) CacheAdd(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.interval)
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > c.interval {
				c.mu.Lock()
				delete(c.Entries, key)
				c.mu.Unlock()
			}
		}
	}
}
