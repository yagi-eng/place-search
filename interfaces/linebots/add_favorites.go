package linebots

import (
	"virtual-travel/usecase"
	"virtual-travel/usecase/dto/userdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// AddFavorites お気に入りリストに追加する
func AddFavorites(UserInteractor usecase.IUserUseCase, bot *linebot.Client, event *linebot.Event, placeID string) {

	lineUserID := event.Source.UserID
	input := userdto.UserCreateInput{
		LineUserID: lineUserID,
	}

	userID := UserInteractor.Create(input)

	res := linebot.NewTextMessage("お気に入りに追加しました！")
	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
