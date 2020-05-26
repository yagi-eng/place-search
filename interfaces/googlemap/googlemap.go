package googlemap

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// GetLocationURLs 検索結果のロケーションのURLを取得する
func GetLocationURLs(c echo.Context, q string) []string {
	locations := searchLocations(c, q)
	placeID := ""
	locationURLs := []string{}

	for i, location := range locations.Results {
		placeID = location.PlaceID
		locationDetail := getLocationDetail(c, placeID)
		locationURLs = append(locationURLs, locationDetail.URL)

		if i+1 == 3 {
			break
		}
	}

	return locationURLs
}

// searchLocations キーワードに基づきロケーションを検索する
func searchLocations(c echo.Context, q string) maps.PlacesSearchResponse {
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

// getLocationDetail ロケーションの詳細情報を取得する
func getLocationDetail(c echo.Context, placeID string) maps.PlaceDetailsResult {
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
