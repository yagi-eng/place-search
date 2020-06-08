package favoritedto

import "github.com/line/line-bot-sdk-go/linebot"

// AddInput DTO
type AddInput struct {
	Bot        *linebot.Client
	ReplyToken string
	LineUserID string
	PlaceID    string
}
