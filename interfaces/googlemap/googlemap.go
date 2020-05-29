package googlemap

import (
	"context"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// PhotoAPIURL Google Maps APIのURL
// SDKでは画像をURL形式で取得できないためAPIを使用
const PhotoAPIURL = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference="

// GetPlaceDetailsAndPhotoURLs キーワードに基づき、プレイスの詳細情報を取得する
func GetPlaceDetailsAndPhotoURLs(c echo.Context, q string) ([]maps.PlaceDetailsResult, []string) {
	placeDetails := []maps.PlaceDetailsResult{}
	placePhotoURLs := []string{}

	places := searchPlaces(c, q)
	for i, place := range places.Results {
		// 最大3件取得
		if i == 3 {
			break
		}

		placeDetail := getPlaceDetail(c, place.PlaceID)
		placeDetails = append(placeDetails, placeDetail)

		photoReference := place.Photos[0].PhotoReference
		placePhotoURL := PhotoAPIURL + photoReference + "&key=" + os.Getenv("GMAP_API_KEY")
		placePhotoURLs = append(placePhotoURLs, placePhotoURL)
	}

	return placeDetails, placePhotoURLs
}

// searchPlaces キーワードに基づき、プレイスを検索する
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

// getPlaceDetail プレイスの詳細情報を取得する
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
