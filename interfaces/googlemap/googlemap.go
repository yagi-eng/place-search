package googlemap

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// GetLocationURL 検索結果のロケーションのURLを取得する
func GetLocationURL(c echo.Context, q string) string {
	searchResult := SearchLocations(c, q)
	placeID := searchResult.Results[0].PlaceID
	locationDetail := GetLocationDetail(c, placeID)
	return locationDetail.URL
}

// SearchLocations キーワードに基づきロケーションを検索する
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

// GetLocationDetail ロケーションの詳細情報を取得する
func GetLocationDetail(c echo.Context, placeID string) maps.PlaceDetailsResult {
	gmc := c.Get("gmc").(*maps.Client)
	r := &maps.PlaceDetailsRequest{
		PlaceID:  placeID,
		Language: "ja",
	}

	res, err := gmc.PlaceDetails(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap PlaceDetails: %v", err)
	}
	return res
}
