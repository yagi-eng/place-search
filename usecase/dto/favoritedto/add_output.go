package favoritedto

import "github.com/line/line-bot-sdk-go/linebot"

// FavoriteAddOutput DTO
type FavoriteAddOutput struct {
	Bot            *linebot.Client
	ReplyToken     string
	IsSuccess      bool
	IsAlreadyAdded bool
}
