package storage

import (
	"fmt"
	"time"
)

type CacheEntry struct {
	value      interface{}
	expiration time.Time
}

type Cache struct {
	data map[string]CacheEntry
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]CacheEntry),
	}
}

func (c *Cache) Set(key string, givenValue interface{}, givenExpiration time.Duration) {
	c.data[key] = CacheEntry{
		value:      givenValue,
		expiration: time.Now().Add(givenExpiration),
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	entry, found := c.data[key]
	if !found || time.Now().After(entry.expiration) {
		return nil, fmt.Errorf("not found")
	}

	return entry.value, nil
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) Update(key string, givenValue interface{}) error {
	entry, found := c.data[key]
	if !found || time.Now().After(entry.expiration) {
		return fmt.Errorf("not found")
	} else {
		c.data[key] = CacheEntry{
			value:      givenValue,
			expiration: entry.expiration,
		}
	}
	return nil
}
