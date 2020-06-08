package favoritedto

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// GetInput DTO
type GetInput struct {
	Bot        *linebot.Client
	ReplyToken string
	LineUserID string
}
