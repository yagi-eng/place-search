package presenter

import (
	"os"
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

// FavoritePresenter お気に入りプレゼンタ
type FavoritePresenter struct {
	bot *linebot.Client
}

// NewFavoritePresenter コンストラクタ
func NewFavoritePresenter() *FavoritePresenter {
	secret := os.Getenv("LBOT_SECRET")
	token := os.Getenv("LBOT_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client: %v", err)
	}

	return &FavoritePresenter{bot: bot}
}

// Add お気に入り追加結果を送信する
func (presenter *FavoritePresenter) Add(out favoritedto.AddOutput) {
	replyToken := out.ReplyToken

	if !out.IsSuccess {
		presenter.replyMessage(msgFail, replyToken)
	} else if out.IsAlreadyAdded {
		presenter.replyMessage(msgAlreadyAdd, replyToken)
	} else {
		presenter.replyMessage(msgSuccess, replyToken)
	}
}

// Get お気に入り一覧を送信する
func (presenter *FavoritePresenter) Get(out favoritedto.GetOutput) {
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
		"お気に入り一覧を表示",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := presenter.bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}

func (presenter *FavoritePresenter) replyMessage(msg string, replyToken string) {
	res := linebot.NewTextMessage(msg)
	if _, err := presenter.bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
