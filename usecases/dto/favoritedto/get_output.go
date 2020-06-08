package favoritedto

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"googlemaps.github.io/maps"
)

// GetOutput DTO
type GetOutput struct {
	Bot            *linebot.Client
	ReplyToken     string
	PlaceDetails   []maps.PlaceDetailsResult
	PlacePhotoURLs []string
}
