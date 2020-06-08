package favoritedto

import (
	"googlemaps.github.io/maps"
)

// GetOutput DTO
type GetOutput struct {
	ReplyToken     string
	PlaceDetails   []maps.PlaceDetailsResult
	PlacePhotoURLs []string
}
