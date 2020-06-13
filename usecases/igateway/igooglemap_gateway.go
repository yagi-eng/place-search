package igateway

import "github.com/yagi-eng/virtual-travel/usecases/dto/googlemapdto"

// IGoogleMapGateway GoogleMapゲートウェイ
type IGoogleMapGateway interface {
	GetPlaceDetailsAndPhotoURLsFromQuery(string) []googlemapdto.Output
	GetPlaceDetailsAndPhotoURLs([]string, bool) []googlemapdto.Output
}
