package controllers

import (
	"context"
	"virtual-travel/util/errmsg"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"googlemaps.github.io/maps"
)

// SearchResult GoogleMapでの検索結果を取得する
func SearchResult() echo.HandlerFunc {
	return func(c echo.Context) error {
		gmc := c.Get("gmc").(*maps.Client)
		r := &maps.TextSearchRequest{
			Query: "東京タワー",
		}

		res, err := gmc.TextSearch(context.Background(), r)
		errmsg.LogFatal(err)

		return c.JSON(fasthttp.StatusOK, res)
	}
}
