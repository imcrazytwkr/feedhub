package groupcache

import (
	"sync"

	"github.com/golang/groupcache/lru"
	"github.com/imcrazytwkr/feedhub/utils/caches"
)

type lruCache[K comparable, V any] struct {
	mu    sync.RWMutex
	cache *lru.Cache
}

func NewLRU[K comparable, V any](size int) (caches.LRU[K, V], error) {
	if size < 1 {
		return nil, ErrInvalidSize
	}

	return &lruCache[K, V]{cache: lru.New(size)}, nil
}

func (c *lruCache[K, V]) Add(key K, value V) {
	c.mu.Lock()
	c.cache.Add(key, value)
	c.mu.Unlock()
}

func (c *lruCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	value, exists := c.cache.Get(key)
	c.mu.Unlock()

	// Do no evil
	if !exists {
		var empty V
		return empty, false
	}

	return value.(V), true
}

func (c *lruCache[K, V]) Remove(key K) {
	c.mu.Lock()
	c.cache.Remove(key)
	c.mu.Unlock()
}

func (c *lruCache[K, V]) RemoveOldest() {
	c.mu.Lock()
	c.cache.RemoveOldest()
	c.mu.Unlock()
}

func (c *lruCache[K, V]) Clear() {
	c.mu.Lock()
	c.cache.Clear()
	c.mu.Unlock()
}

func (c *lruCache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache.Len()
}
