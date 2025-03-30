package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]caheEntry
	mu   sync.RWMutex
	ttl  time.Duration
}

type caheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = caheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.Data[key]
	return data.val, ok
}

func (c *Cache) cleanCache(cutoff time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, data := range c.Data {
		if data.createdAt.Before(cutoff) {
			delete(c.Data, key)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for t := range ticker.C {
		c.cleanCache(t.Add(-c.ttl))
	}
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		Data: make(map[string]caheEntry),
		ttl:  ttl,
	}
	go cache.reapLoop()
	return cache
}
