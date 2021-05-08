package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type API struct {
	cache *Cache
}

func Start() {
	cache := NewCache()
	api := &API{cache: cache}

	e := echo.New()
	e.GET("/:key", api.handleGet)
	e.POST("/:key", api.handleSet)
	e.DELETE("/:key", api.handleDelete)

	e.Logger.Fatal(e.Start(":1337"))
}

func (api API) handleGet(c echo.Context) error {
	key := c.Param("key")

	value, err := api.cache.Get(key)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("[%s]", err.Error()))
	}

	return c.String(http.StatusOK, value)
}

func (api API) handleSet(c echo.Context) error {
	key := c.Param("key")
	value := c.FormValue("value")

	api.cache.Set(key, value)
	return nil
}

func (api API) handleDelete(c echo.Context) error {
	key := c.Param("key")

	api.cache.Delete(key)
	return nil
}
