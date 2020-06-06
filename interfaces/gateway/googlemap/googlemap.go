package googlemap

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// maxDetails プレイスの詳細情報の最大取得件数
const maxDetails = 3

// PhotoAPIURL Google Maps APIのURL
// SDKでは画像をURL形式で取得できないためAPIを使用
const PhotoAPIURL = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference="

// GetPlaceDetailsAndPhotoURLs キーワードに基づき、プレイスの詳細情報を取得する
func GetPlaceDetailsAndPhotoURLs(gm *maps.Client, q string) ([]maps.PlaceDetailsResult, []string) {
	placeDetails := []maps.PlaceDetailsResult{}
	placePhotoURLs := []string{}

	places := searchPlaces(gm, q)
	for i, place := range places.Results {
		if i == maxDetails {
			break
		}

		placeDetail := getPlaceDetail(gm, place.PlaceID)
		placeDetails = append(placeDetails, placeDetail)

		placePhotoURL := getPlacePhotoURL(place.Photos[0].PhotoReference)
		placePhotoURLs = append(placePhotoURLs, placePhotoURL)
	}

	return placeDetails, placePhotoURLs
}

// searchPlaces キーワードに基づき、プレイスを検索する
// 単独での使用を想定して第一引数には echo.Context を渡す
func searchPlaces(gm *maps.Client, q string) maps.PlacesSearchResponse {
	r := &maps.TextSearchRequest{
		Query:    q,
		Language: "ja",
		Location: &maps.LatLng{Lat: 35.658517, Lng: 139.70133399999997}, // 渋谷
		Radius:   50000,
	}

	res, err := gm.TextSearch(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap TextSearch: %v", err)
	}
	return res
}

// getPlaceDetail プレイスの詳細情報を取得する
// 単独での使用を想定して第一引数には echo.Context を渡す
func getPlaceDetail(gm *maps.Client, placeID string) maps.PlaceDetailsResult {
	r := &maps.PlaceDetailsRequest{
		PlaceID:  placeID,
		Language: "ja",
	}

	res, err := gm.PlaceDetails(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap PlaceDetails: %v", err)
	}
	return res
}

// getPlacePhotoURL プレイスの写真のURLを取得する
func getPlacePhotoURL(photoReference string) string {
	return PhotoAPIURL + photoReference + "&key=" + os.Getenv("GMAP_API_KEY")
}
