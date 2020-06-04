package linebots

import (
	"virtual-travel/usecase"
	"virtual-travel/usecase/dto/userdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// AddFavorites お気に入りリストに追加する
func AddFavorites(interactor usecase.IUserUseCase, bot *linebot.Client, event *linebot.Event) {
	LineUserID := event.Source.UserID
	input := userdto.UserCreateInput{
		LineUserID: LineUserID,
	}

	interactor.Create(input)

	res := linebot.NewTextMessage("お気に入りに追加しました！（してない）")
	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
