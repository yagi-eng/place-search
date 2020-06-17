package gateway

import (
	"context"
	"os"

	"github.com/yagi-eng/place-search/usecases/dto/googlemapdto"

	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// GoogleMapGateway GoogleMapゲートウェイ
type GoogleMapGateway struct {
	gmc *maps.Client
}

// NewGoogleMapGateway コンストラクタ
func NewGoogleMapGateway() *GoogleMapGateway {
	apiKey := os.Getenv("GMAP_API_KEY")
	gmc, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		logrus.Fatalf("Error creating GoogleMap client: %v", err)
	}
	return &GoogleMapGateway{gmc: gmc}
}

// maxDetailsOfSearch 検索結果の最大取得件数
const maxDetailsOfSearch = 4

// maxDetailsOfFavorite お気に入り一覧の最大表示件数
const maxDetailsOfFavorite = 10

// photoAPIURL Google Maps APIのURL
// SDKでは画像をURL形式で取得できないためAPIを使用
const photoAPIURL = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference="
const noImageURL = "https://1.bp.blogspot.com/-D2I7Z7-HLGU/Xlyf7OYUi8I/AAAAAAABXq4/jZ0035aDGiE5dP3WiYhlSqhhMgGy8p7zACNcBGAsYHQ/s1600/no_image_square.jpg"

/*****
*
* データ成型部分
*
******/

// GetPlaceDetailsAndPhotoURLsWithQuery キーワードに基づき、プレイスの詳細情報を取得する
func (gateway *GoogleMapGateway) GetPlaceDetailsAndPhotoURLsWithQuery(q string) []googlemapdto.Output {
	places := gateway.searchPlacesWithQuery(q)
	placeIDs := gateway.getPlaceIDs(places.Results)

	return gateway.GetPlaceDetailsAndPhotoURLs(placeIDs, false)
}

// GetPlaceDetailsAndPhotoURLsWithQueryLatLng キーワード、経度/緯度に基づき、プレイスの詳細情報を取得する
func (gateway *GoogleMapGateway) GetPlaceDetailsAndPhotoURLsWithQueryLatLng(q string, lat float64, lng float64) []googlemapdto.Output {
	places := gateway.searchPlacesWithLatLng(q, lat, lng)
	placeIDs := gateway.getPlaceIDs(places.Results)

	return gateway.GetPlaceDetailsAndPhotoURLs(placeIDs, false)
}

// GetPlaceDetailsAndPhotoURLs placeIDsに基づき、プレイスの詳細情報を取得する
func (gateway *GoogleMapGateway) GetPlaceDetailsAndPhotoURLs(placeIDs []string, isFavorite bool) []googlemapdto.Output {
	googleMapOutputs := []googlemapdto.Output{}

	maxDetails := maxDetailsOfSearch
	if isFavorite {
		maxDetails = maxDetailsOfFavorite
	}

	for i, placeID := range placeIDs {
		if i == maxDetails {
			break
		}

		placeDetail := gateway.getPlaceDetail(placeID)
		placePhotoURL := noImageURL
		if len(placeDetail.Photos) > 0 {
			placePhotoURL = gateway.getPlacePhotoURL(placeDetail.Photos[0].PhotoReference)
		}

		googleMapOutput := googlemapdto.Output{
			Name:     placeDetail.Name,
			PlaceID:  placeDetail.PlaceID,
			Address:  placeDetail.FormattedAddress,
			URL:      placeDetail.URL,
			PhotoURL: placePhotoURL,
		}

		googleMapOutputs = append(googleMapOutputs, googleMapOutput)
	}

	return googleMapOutputs
}

// getPlaceIDs プレイスの検索結果からplaceIDを取得する
func (gateway *GoogleMapGateway) getPlaceIDs(places []maps.PlacesSearchResult) []string {
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

// searchPlacesWithQuery キーワードに基づき、プレイスを検索する
func (gateway *GoogleMapGateway) searchPlacesWithQuery(q string) maps.PlacesSearchResponse {
	r := &maps.TextSearchRequest{
		Query:    q,
		Language: "ja",
		Location: &maps.LatLng{Lat: 35.658517, Lng: 139.70133399999997}, // 渋谷
		Radius:   10,
	}

	res, err := gateway.gmc.TextSearch(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap TextSearch: %v", err)
	}
	return res
}

// searchPlacesWithLatLng キーワード、経度/緯度に基づき、プレイスを検索する
func (gateway *GoogleMapGateway) searchPlacesWithLatLng(q string, lat float64, lng float64) maps.PlacesSearchResponse {
	r := &maps.TextSearchRequest{
		Query:    q,
		Language: "ja",
		Location: &maps.LatLng{Lat: lat, Lng: lng},
		Radius:   500,
	}

	res, err := gateway.gmc.TextSearch(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap TextSearch: %v", err)
	}
	return res
}

// getPlaceDetail プレイスの詳細情報を取得する
func (gateway *GoogleMapGateway) getPlaceDetail(placeID string) maps.PlaceDetailsResult {
	nameFM, _ := maps.ParsePlaceDetailsFieldMask("name")
	placeIDFM, _ := maps.ParsePlaceDetailsFieldMask("place_id")
	addrFM, _ := maps.ParsePlaceDetailsFieldMask("formatted_address")
	urlFM, _ := maps.ParsePlaceDetailsFieldMask("url")
	photoFM, _ := maps.ParsePlaceDetailsFieldMask("photos")

	r := &maps.PlaceDetailsRequest{
		PlaceID:  placeID,
		Language: "ja",
		Fields: []maps.PlaceDetailsFieldMask{
			nameFM,
			placeIDFM,
			addrFM,
			urlFM,
			photoFM,
		},
	}

	res, err := gateway.gmc.PlaceDetails(context.Background(), r)
	if err != nil {
		logrus.Fatal("Error GoogleMap PlaceDetails: %v", err)
	}
	return res
}

// getPlacePhotoURL プレイスの写真のURLを取得する
func (gateway *GoogleMapGateway) getPlacePhotoURL(photoReference string) string {
	return photoAPIURL + photoReference + "&key=" + os.Getenv("GMAP_API_KEY")
}
