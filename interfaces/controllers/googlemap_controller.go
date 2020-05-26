package controllers

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
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
		if err != nil {
			logrus.Fatal("Error GoogleMap TextSearch: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
