package linebots

import (
	"virtual-travel/usecase"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

var uc usecase.IUserCreateUseCase

// AddFavorites お気に入りリストに追加する
func AddFavorites(c echo.Context, bot *linebot.Client, event *linebot.Event) {
	LineUserID := event.Source.UserID
	input := usecase.UserCreateInput{
		LineUserID: LineUserID,
	}

	uc.Handle(input)

	res := linebot.NewTextMessage("お気に入りに追加しました！（してない）")
	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
