package searchdto

import "googlemaps.github.io/maps"

// Output DTO
type Output struct {
	ReplyToken     string
	Q              string
	PlaceDetails   []maps.PlaceDetailsResult
	PlacePhotoURLs []string
}
