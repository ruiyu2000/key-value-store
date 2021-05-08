package server

import "errors"

type Cache struct {
	cache map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Get(key string) (value interface{}, err error) {
	value, ok := c.cache[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return value, nil
}

func (c Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}
