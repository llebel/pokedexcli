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
func (c *Cache) reapLoop(interval time.Duration) {
	// Initialize reap ticker
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

// reap -
func (c *Cache) reap(now time.Time, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entryMap {
		// If entry is older than cache TTL
		if now.After(entry.createdAt.Add(ttl)) {
			delete(c.entryMap, key)
		}
	}
}
