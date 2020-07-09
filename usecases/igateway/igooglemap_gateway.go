package igateway

import "github.com/yagi-eng/place-search/domain/model"

// IGoogleMapGateway GoogleMapゲートウェイ
type IGoogleMapGateway interface {
	GetPlaceDetailsAndPhotoURLsWithQuery(string) []model.Place
	GetPlaceDetailsAndPhotoURLsWithQueryLatLng(string, float64, float64) []model.Place
	GetPlaceDetailsAndPhotoURLs([]string, bool) []model.Place
}
