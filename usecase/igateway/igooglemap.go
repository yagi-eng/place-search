package igateway

import "googlemaps.github.io/maps"

// IGoogleMapGateway Googleゲートウェイ
type IGoogleMapGateway interface {
	GetPlaceDetailsAndPhotoURLsFromQuery(string) ([]maps.PlaceDetailsResult, []string)
	GetPlaceDetailsAndPhotoURLs([]string, bool) ([]maps.PlaceDetailsResult, []string)
}
