package linebots

import (
	"virtual-travel/usecase"
	"virtual-travel/usecase/dto/favoritedto"
	"virtual-travel/usecase/dto/userdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

const msgNotAdd = "お気に入りに追加できませんでした。再度登録をお願いしますm(__)m"
const msgAlreadyAdd = "既に追加済みです！"
const msgAdd = "お気に入りに追加しました！"

// AddFavorites お気に入りリストに追加する
func AddFavorites(userInteractor usecase.IUserUseCase, favoriteInteractor usecase.IFavoriteUseCase,
	bot *linebot.Client, event *linebot.Event, placeID string) {

	lineUserID := event.Source.UserID
	userCreateInput := userdto.UserCreateInput{
		LineUserID: lineUserID,
	}

	userCreateOutput := userInteractor.Create(userCreateInput)
	userID := userCreateOutput.UserID

	if userID == 0 {
		replyMessage(bot, msgNotAdd, event.ReplyToken)
		return
	}

	favoriteAddInput := favoritedto.FavoriteAddInput{
		UserID:  userID,
		PlaceID: placeID,
	}

	favoriteAddOutput := favoriteInteractor.Add(favoriteAddInput)
	isAlreadyAdded := favoriteAddOutput.IsAlreadyAdded

	if isAlreadyAdded {
		replyMessage(bot, msgAlreadyAdd, event.ReplyToken)
		return
	}

	replyMessage(bot, msgAdd, event.ReplyToken)
}

func replyMessage(bot *linebot.Client, msg string, replyToken string) {
	res := linebot.NewTextMessage(msg)
	if _, err := bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
