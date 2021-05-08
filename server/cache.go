package server

import "errors"

type Cache struct {
	cache map[string]string
}

var ErrNotFound = errors.New("not found")

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]string),
	}
}

func (c Cache) Get(key string) (value string, err error) {
	value, ok := c.cache[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (c Cache) Set(key string, value string) {
	c.cache[key] = value
}

func (c Cache) Delete(key string) {
	delete(c.cache, key)
}
