package presenter

import (
	"os"
	"unicode/utf8"
	"virtual-travel/usecases/dto/searchdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

const msgNoResult = "検索結果は0件でした"

// SearchPresenter 検索プレゼンタ
type SearchPresenter struct {
	bot *linebot.Client
}

// NewSearchPresenter コンストラクタ
func NewSearchPresenter() *SearchPresenter {
	secret := os.Getenv("LBOT_SECRET")
	token := os.Getenv("LBOT_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client: %v", err)
	}

	return &SearchPresenter{bot: bot}
}

// Hundle 検索結果を送信する
func (presenter *SearchPresenter) Hundle(out searchdto.Output) {
	replyToken := out.ReplyToken
	placeDetails := out.PlaceDetails

	if len(placeDetails) == 0 {
		presenter.replyMessage(msgSuccess, replyToken)
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
		"「"+out.Q+"」の検索結果です",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := presenter.bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}

func (presenter *SearchPresenter) replyMessage(msg string, replyToken string) {
	res := linebot.NewTextMessage(msg)
	if _, err := presenter.bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
