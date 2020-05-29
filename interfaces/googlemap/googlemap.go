package googlemap

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// GetPlaceDetails 検索結果のロケーションのURLを取得する
func GetPlaceDetails(c echo.Context, q string) []maps.PlaceDetailsResult {
	placeDetails := []maps.PlaceDetailsResult{}

	places := searchPlaces(c, q)
	for i, place := range places.Results {
		placeID := place.PlaceID
		placeDetail := getPlaceDetail(c, placeID)
		placeDetails = append(placeDetails, placeDetail)

		if i+1 == 3 {
			break
		}
	}

	return placeDetails
}

// searchPlaces キーワードに基づきロケーションを検索する
// 単独での使用を想定して第一引数には echo.Context を渡す
func searchPlaces(c echo.Context, q string) maps.PlacesSearchResponse {
	gmc := c.Get("gmc").(*maps.Client)

	r := &maps.TextSearchRequest{
		Query:    q,
		Language: "ja",
		Location: &maps.LatLng{Lat: 35.658517, Lng: 139.70133399999997}, // 渋谷
		Radius:   50000,
	}

	res, err := gmc.TextSearch(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap TextSearch: %v", err)
	}
	return res
}

// getPlaceDetail ロケーションの詳細情報を取得する
// 単独での使用を想定して第一引数には echo.Context を渡す
func getPlaceDetail(c echo.Context, placeID string) maps.PlaceDetailsResult {
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
