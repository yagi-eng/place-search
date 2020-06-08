package favoritedto

import "github.com/line/line-bot-sdk-go/linebot"

// AddOutput DTO
type AddOutput struct {
	Bot            *linebot.Client
	ReplyToken     string
	IsSuccess      bool
	IsAlreadyAdded bool
}
