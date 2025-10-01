package pokecache

import (
	"time"
)

// Add -
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entryMap[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if entry, exist := c.entryMap[key]; exist {
		return entry.val, exist
	}
	return nil, false
}

// reapLoop -
func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.entryMap {
		// If entry is older than cache TTL
		if now.After(entry.createdAt.Add(c.ttl)) {
			delete(c.entryMap, key)
		}
	}
}
