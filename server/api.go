package server

import (
	"github.com/labstack/echo/v4"
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
	return nil
}

func (api API) handleSet(c echo.Context) error {
	return nil
}
