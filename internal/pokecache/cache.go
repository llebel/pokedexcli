package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entryMap map[string]cacheEntry
	mu       *sync.RWMutex
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
	}

	go c.reapLoop(interval)

	return c
}
