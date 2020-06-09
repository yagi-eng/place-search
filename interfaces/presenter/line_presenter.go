package presenter

import (
	"os"
	"unicode/utf8"
	"virtual-travel/usecases/dto/favoritedto"
	"virtual-travel/usecases/dto/searchdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// AddFavoriteで使用
const msgFail = "お気に入りに追加できませんでした。再度登録をお願いしますm(__)m"
const msgAlreadyAdd = "既に追加済みです！"
const msgSuccess = "お気に入りに追加しました！"

// GetFavoriteで使用
const msgNoRegistGF = "お気に入り登録されていません"
const msgAltTextGF = "お気に入りを一覧表示しました"

// Searchで使用
const msgNoRegistS = "検索結果は0件でした"

const maxTextWC = 60

// LinePresenter LINEプレゼンタ
type LinePresenter struct {
	bot *linebot.Client
}

type carouselMsgs struct {
	noResult string
	altText  string
}

// NewLinePresenter コンストラクタ
func NewLinePresenter() *LinePresenter {
	secret := os.Getenv("LBOT_SECRET")
	token := os.Getenv("LBOT_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client: %v", err)
	}

	return &LinePresenter{bot: bot}
}

// AddFavorite お気に入り追加結果を送信する
func (presenter *LinePresenter) AddFavorite(out favoritedto.AddOutput) {
	replyToken := out.ReplyToken

	if !out.IsSuccess {
		presenter.replyMessage(msgFail, replyToken)
	} else if out.IsAlreadyAdded {
		presenter.replyMessage(msgAlreadyAdd, replyToken)
	} else {
		presenter.replyMessage(msgSuccess, replyToken)
	}
}

// GetFavorite お気に入り一覧を送信する
func (presenter *LinePresenter) GetFavorite(out favoritedto.GetOutput) {
	msgs := carouselMsgs{
		noResult: msgNoRegistGF,
		altText:  msgAltTextGF,
	}
	presenter.replyCarouselColumn(msgs, out.PlaceDetails, out.PlacePhotoURLs, out.ReplyToken)
}

// Search 検索結果を送信する
func (presenter *LinePresenter) Search(out searchdto.Output) {
	msgs := carouselMsgs{
		noResult: msgNoRegistS,
		altText:  "「" + out.Q + "」の検索結果です",
	}
	presenter.replyCarouselColumn(msgs, out.PlaceDetails, out.PlacePhotoURLs, out.ReplyToken)
}

func (presenter *LinePresenter) replyCarouselColumn(msgs carouselMsgs,
	placeDetails []maps.PlaceDetailsResult, placePhotoURLs []string, replyToken string) {

	if len(placeDetails) == 0 {
		presenter.replyMessage(msgs.noResult, replyToken)
		return
	}

	ccs := []*linebot.CarouselColumn{}
	for i, pd := range placeDetails {
		formattedAddress := pd.FormattedAddress
		if maxTextWC < utf8.RuneCountInString(pd.FormattedAddress) {
			formattedAddress = string([]rune(pd.FormattedAddress)[:maxTextWC])
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
		msgs.altText,
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := presenter.bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}

func (presenter *LinePresenter) replyMessage(msg string, replyToken string) {
	res := linebot.NewTextMessage(msg)
	if _, err := presenter.bot.ReplyMessage(replyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
