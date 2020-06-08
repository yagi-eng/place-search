package presenter

import (
	"unicode/utf8"
	"virtual-travel/usecases/dto/favoritedto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

const msgFail = "お気に入りに追加できませんでした。再度登録をお願いしますm(__)m"
const msgAlreadyAdd = "既に追加済みです！"
const msgSuccess = "お気に入りに追加しました！"
const msgNoRegist = "お気に入り登録されていません"

const maxTextWC = 60

// Add お気に入り追加結果を送信する
func Add(out favoritedto.AddOutput) {
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

// Get お気に入り一覧を送信する
func Get(out favoritedto.GetOutput) {
	bot := out.Bot
	replyToken := out.ReplyToken
	placeDetails := out.PlaceDetails

	if len(placeDetails) == 0 {
		replyMessage(bot, msgSuccess, replyToken)
		return
	}

	ccs := []*linebot.CarouselColumn{}
	for i, pd := range placeDetails {
		formattedAddress := pd.FormattedAddress
		if maxTextWC < utf8.RuneCountInString(pd.FormattedAddress) {
			formattedAddress = string([]rune(pd.FormattedAddress)[:maxTextWC])
		}

		cc := linebot.NewCarouselColumn(
			out.PlacePhotoURLs[i],
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

	if _, err := bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}

func replyMessage(bot *linebot.Client, msg string, replyToken string) {
	res := linebot.NewTextMessage(msg)
	if _, err := bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
