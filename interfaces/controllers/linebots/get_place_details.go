// package linebots

// import (
// 	"unicode/utf8"
// 	"virtual-travel/interfaces/gateway/googlemap"

// 	"github.com/line/line-bot-sdk-go/linebot"
// 	"github.com/sirupsen/logrus"
// 	"googlemaps.github.io/maps"
// )

// const spot = " 観光地"

// // GetPlaceDetails プレイスの詳細情報を取得して応答する
// func GetPlaceDetails(gm *maps.Client, bot *linebot.Client, event *linebot.Event, q string) {
// 	placeDetails, placePhotoURLs := googlemap.GetPlaceDetailsAndPhotoURLsFromQuery(gm, q+spot)

// 	if len(placeDetails) == 0 {
// 		res := linebot.NewTextMessage("検索結果は0件でした")
// 		if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
// 			logrus.Fatalf("Error LINEBOT replying message: %v", err)
// 		}
// 		return
// 	}

// 	ccs := []*linebot.CarouselColumn{}
// 	for i, pd := range placeDetails {
// 		formattedAddress := pd.FormattedAddress
// 		if 60 < utf8.RuneCountInString(pd.FormattedAddress) {
// 			formattedAddress = string([]rune(pd.FormattedAddress)[:60])
// 		}

// 		cc := linebot.NewCarouselColumn(
// 			placePhotoURLs[i],
// 			pd.Name,
// 			formattedAddress,
// 			linebot.NewURIAction("Open Google Map", pd.URL),
// 			linebot.NewPostbackAction("Add to my favorites", "action=favorite&placeId="+pd.PlaceID, "", ""),
// 		).WithImageOptions("#FFFFFF")
// 		ccs = append(ccs, cc)
// 	}

// 	res := linebot.NewTemplateMessage(
// 		"「"+q+"」の検索結果です",
// 		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
// 	)

// 	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
// 		logrus.Fatalf("Error LINEBOT replying message: %v", err)
// 	}
// }
