package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entryMap map[string]cacheEntry
	mu       *sync.RWMutex
	ttl      time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		entryMap: map[string]cacheEntry{},
		mu:       &sync.RWMutex{},
		ttl:      interval,
	}

	// Initialize reap ticker
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reapLoop()
		}
	}()

	return c
}
