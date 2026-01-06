package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	entries		map[string]cacheEntry
	mu			sync.Mutex
	interval 	time.Duration
}

func NewCache(interval time.Duration) *Cache {
	retCache := Cache {
		entries:	make(map[string]cacheEntry),
		interval:	interval,
	}
	go retCache.reapLoop()

	return &retCache
}

func (c *Cache) Add(key string, value []byte) {
	// add the key and value
	//
	// remember to lock access using the Mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry {
		createdAt:	time.Now(),
		val:		value,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// get an entry from the cache
	// -> true if found, false if not found
	//
	// remember to lock access using the Mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	// each time an interval passes, remove older entries
	//
	// remember to lock access using the Mutex
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

