package favoritedto

import "github.com/line/line-bot-sdk-go/linebot"

// FavoriteAddInput DTO
type FavoriteAddInput struct {
	Bot        *linebot.Client
	ReplyToken string
	LineUserID string
	PlaceID    string
}
