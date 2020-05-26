package controllers

import (
	"virtual-travel/interfaces/googlemap"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

// SearchLocations GoogleMapでの検索結果を取得する
func SearchLocations() echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("q")
		res := googlemap.SearchLocations(c, q)
		return c.JSON(fasthttp.StatusOK, res)
	}
}
