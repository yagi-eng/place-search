package linebots

import (
	"virtual-travel/usecase"
	"virtual-travel/usecase/dto/favoritedto"
	"virtual-travel/usecase/dto/userdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// AddFavorites お気に入りリストに追加する
func AddFavorites(UserInteractor usecase.IUserUseCase, FavoriteInteractor usecase.IFavoriteUseCase,
	bot *linebot.Client, event *linebot.Event, placeID string) {

	lineUserID := event.Source.UserID
	userCreateInput := userdto.UserCreateInput{
		LineUserID: lineUserID,
	}

	userCreateOutput := UserInteractor.Create(userCreateInput)

	favoriteAddInput := favoritedto.FavoriteAddInput{
		UserID:  userCreateOutput.UserID,
		PlaceID: placeID,
	}
	FavoriteInteractor.Add(favoriteAddInput)

	res := linebot.NewTextMessage("お気に入りに追加しました！")
	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
