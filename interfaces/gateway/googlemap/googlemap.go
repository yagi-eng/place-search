package googlemap

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// maxDetailsOfSearch 検索結果の最大取得件数
const maxDetailsOfSearch = 3

// maxDetailsOfFavorite お気に入り一覧の最大表示件数
const maxDetailsOfFavorite = 3

// photoAPIURL Google Maps APIのURL
// SDKでは画像をURL形式で取得できないためAPIを使用
const photoAPIURL = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference="

/*****
*
* データ成型部分
*
******/

// GetPlaceDetailsAndPhotoURLsFromQuery キーワードに基づき、プレイスの詳細情報を取得する
func GetPlaceDetailsAndPhotoURLsFromQuery(gm *maps.Client, q string) ([]maps.PlaceDetailsResult, []string) {
	places := searchPlaces(gm, q)
	placeIDs := getPlaceIDs(places.Results)

	return GetPlaceDetailsAndPhotoURLs(gm, placeIDs, false)
}

// GetPlaceDetailsAndPhotoURLs キーワードに基づき、プレイスの詳細情報を取得する
func GetPlaceDetailsAndPhotoURLs(gm *maps.Client, placeIDs []string, isFavorite bool) ([]maps.PlaceDetailsResult, []string) {
	placeDetails := []maps.PlaceDetailsResult{}
	placePhotoURLs := []string{}

	maxDetails := maxDetailsOfSearch
	if isFavorite {
		maxDetails = maxDetailsOfFavorite
	}

	for i, placeID := range placeIDs {
		if i == maxDetails {
			break
		}

		placeDetail := getPlaceDetail(gm, placeID)
		placeDetails = append(placeDetails, placeDetail)

		placePhotoURL := getPlacePhotoURL(placeDetail.Photos[0].PhotoReference)
		placePhotoURLs = append(placePhotoURLs, placePhotoURL)
	}

	return placeDetails, placePhotoURLs
}

// getPlaceIDs プレイスの検索結果からplaceIDを取得する
func getPlaceIDs(places []maps.PlacesSearchResult) []string {
	placeIDs := []string{}
	for _, place := range places {
		placeIDs = append(placeIDs, place.PlaceID)
	}
	return placeIDs
}

/*****
*
* 通信部分
*
******/

// searchPlaces キーワードに基づき、プレイスを検索する
// 単独での使用を想定して第一引数には *maps.Client を渡す
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
// 単独での使用を想定して第一引数には *maps.Client を渡す
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
	return photoAPIURL + photoReference + "&key=" + os.Getenv("GMAP_API_KEY")
}
