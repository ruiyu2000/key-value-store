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
	e.Logger.Fatal(e.Start(":1337"))
}

func (api API) handleGet(c echo.Context) error {
	key := c.Param("key")

	value, err := api.cache.Get(key)
	if err != nil {
		return nil
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", value))
}

func (api API) handleSet(c echo.Context) error {
	key := c.Param("key")
	value := c.FormValue("value")

	api.cache.Set(key, value)
	return nil
}
