package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cache map[string]cacheEntry
	mux sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache{
	new := Cache{
		cache: make(map[string]cacheEntry),
	}
	go new.reapLoop(interval)
	return new
}

func (c *Cache) Add(key string, value []byte){
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	value, ok := c.cache[key]
	if !ok{
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop(inter time.Duration){
	loop := time.NewTicker(inter)
	for range loop.C{
		c.reap(inter)
	}
}

func (c *Cache) reap(inter time.Duration){
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, cacheInfo := range c.cache{
		totalTime := time.Now().Add(inter)
		if cacheInfo.createdAt.Before(totalTime){
			delete(c.cache, key)
		}
	}
}