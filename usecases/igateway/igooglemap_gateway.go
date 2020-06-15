package igateway

import "github.com/yagi-eng/place-search/usecases/dto/googlemapdto"

// IGoogleMapGateway GoogleMapゲートウェイ
type IGoogleMapGateway interface {
	GetPlaceDetailsAndPhotoURLsFromQuery(string) []googlemapdto.Output
	GetPlaceDetailsAndPhotoURLs([]string, bool) []googlemapdto.Output
}
