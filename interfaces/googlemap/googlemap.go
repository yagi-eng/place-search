package googlemap

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// SearchLocations GoogleMapの検索結果を取得する
func SearchLocations(c echo.Context, q string) maps.PlacesSearchResponse {
	gmc := c.Get("gmc").(*maps.Client)
	r := &maps.TextSearchRequest{
		Query:    q,
		Language: "ja",
	}

	res, err := gmc.TextSearch(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap TextSearch: %v", err)
	}
	return res
}
