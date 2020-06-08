package presenter

import (
	"virtual-travel/usecase/dto/favoritedto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

const msgFail = "お気に入りに追加できませんでした。再度登録をお願いしますm(__)m"
const msgAlreadyAdd = "既に追加済みです！"
const msgSuccess = "お気に入りに追加しました！"

// Add お気に入りプレゼンター
func Add(out favoritedto.FavoriteAddOutput) {
	bot := out.Bot
	replyToken := out.ReplyToken

	if !out.IsSuccess {
		replyMessage(bot, msgFail, replyToken)
	} else if out.IsAlreadyAdded {
		replyMessage(bot, msgAlreadyAdd, replyToken)
	} else {
		replyMessage(bot, msgSuccess, replyToken)
	}
}

func replyMessage(bot *linebot.Client, msg string, replyToken string) {
	res := linebot.NewTextMessage(msg)
	if _, err := bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
