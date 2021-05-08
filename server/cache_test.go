package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache(t *testing.T) {
	key := "somekey"
	val := "asdf12345"

	cache := NewCache()
	_, err := cache.Get(key)
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)

	cache.Set(key, val)
	actual, err := cache.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, actual)

	cache.Delete(key)
	_, err = cache.Get(key)
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}
