package botreply

import (
	"virtual-travel/interfaces/googlemap"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// GetPlaceDetails プレイスの詳細情報を取得して応答する
func GetPlaceDetails(c echo.Context, bot *linebot.Client, event *linebot.Event, q string) {
	placeDetails, placePhotoURLs := googlemap.GetPlaceDetailsAndPhotoURLs(c, q)

	if len(placeDetails) == 0 {
		res := linebot.NewTextMessage("検索結果は0件でした")
		if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
			logrus.Fatalf("Error LINEBOT replying message: %v", err)
		}
		return
	}

	ccs := []*linebot.CarouselColumn{}
	for i, pd := range placeDetails {
		cc := linebot.NewCarouselColumn(
			placePhotoURLs[i],
			pd.Name,
			pd.FormattedAddress,
			linebot.NewURIAction("Open Google Map", pd.URL),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}

	res := linebot.NewTemplateMessage(
		"「"+q+"」の検索結果です",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
