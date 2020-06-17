package igateway

import "github.com/yagi-eng/place-search/usecases/dto/googlemapdto"

// IGoogleMapGateway GoogleMapゲートウェイ
type IGoogleMapGateway interface {
	GetPlaceDetailsAndPhotoURLsWithQuery(string) []googlemapdto.Output
	GetPlaceDetailsAndPhotoURLsWithQueryLatLng(string, float64, float64) []googlemapdto.Output
	GetPlaceDetailsAndPhotoURLs([]string, bool) []googlemapdto.Output
}
