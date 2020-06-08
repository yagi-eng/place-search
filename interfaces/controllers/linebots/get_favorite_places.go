package linebots

import (
	"unicode/utf8"
	"virtual-travel/interfaces/gateway/googlemap"
	"virtual-travel/usecase"
	"virtual-travel/usecase/dto/favoritedto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// GetFavoritePlaces お気に入り登録されたプレイスを表示する
func GetFavoritePlaces(favoriteInteractor usecase.IFavoriteUseCase, gm *maps.Client, bot *linebot.Client, event *linebot.Event) {
	lineUserID := event.Source.UserID
	favoriteGetInput := favoritedto.FavoriteGetInput{
		LineUserID: lineUserID,
	}

	favoriteGetOutput := favoriteInteractor.Get(favoriteGetInput)
	PlaceIDs := favoriteGetOutput.PlaceIDs

	placeDetails, placePhotoURLs := googlemap.GetPlaceDetailsAndPhotoURLs(gm, PlaceIDs, true)

	if len(placeDetails) == 0 {
		res := linebot.NewTextMessage("お気に入り登録されていません")
		if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
			logrus.Fatalf("Error LINEBOT replying message: %v", err)
		}
		return
	}

	ccs := []*linebot.CarouselColumn{}
	for i, pd := range placeDetails {
		formattedAddress := pd.FormattedAddress
		if 60 < utf8.RuneCountInString(pd.FormattedAddress) {
			formattedAddress = string([]rune(pd.FormattedAddress)[:60])
		}

		cc := linebot.NewCarouselColumn(
			placePhotoURLs[i],
			pd.Name,
			formattedAddress,
			linebot.NewURIAction("Open Google Map", pd.URL),
			linebot.NewPostbackAction("Add to my favorites", "action=favorite&placeId="+pd.PlaceID, "", ""),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}

	res := linebot.NewTemplateMessage(
		"お気に入り一覧を表示",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
