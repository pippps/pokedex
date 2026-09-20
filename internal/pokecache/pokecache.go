package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMapEntry map[string]cacheEntry
	mux           sync.RWMutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMapEntry: make(map[string]cacheEntry),
		mux:           sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cacheMapEntry[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	val, ok := c.cacheMapEntry[key]
	c.mux.RUnlock()

	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cacheMapEntry {
		age := time.Since(v.createdAt)
		if age > interval {
			delete(c.cacheMapEntry, k)
		}
	}
}
