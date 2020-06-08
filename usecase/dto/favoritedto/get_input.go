package favoritedto

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"googlemaps.github.io/maps"
)

// GetInput DTO
type GetInput struct {
	Gmc        *maps.Client
	Bot        *linebot.Client
	ReplyToken string
	LineUserID string
}
