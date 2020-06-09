package presenter

import (
	"fmt"
	"os"
	"unicode/utf8"
	"virtual-travel/usecases/dto/favoritedto"
	"virtual-travel/usecases/dto/searchdto"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// AddFavoriteで使用
const msgFailAF = "お気に入りに追加できませんでした。再度登録をお願いしますm(__)m"
const msgAlreadyAddAF = "既に追加済みです！"
const msgSuccessAF = "お気に入りに追加しました！"

// GetFavoritesで使用
const msgNoRegistGF = "お気に入り登録されていません"
const msgAltTextGF = "お気に入り一覧の表示結果です"
const msgPostbackActionLabelGF = "Remove"
const msgPostbackActionDataGF = "action=removeFavorite&placeId=%s"

// RemoveFavoriteで使用
const msgFailRF = "お気に入りを削除できませんでした。再度登録をお願いしますm(__)m"
const msgAlreadyAddRF = "既に削除済みです！"
const msgSuccessRF = "お気に入りを削除しました！"

// Searchで使用
const msgNoRegistS = "検索結果は0件でした"
const msgAltTextS = "「%s」の検索結果です"
const msgPostbackActionLabelS = "Add to my favorites"
const msgPostbackActionDataS = "action=addFavorite&placeId=%s"

const maxTextWC = 60

// LinePresenter LINEプレゼンタ
type LinePresenter struct {
	bot *linebot.Client
}

type carouselMsgs struct {
	noResult            string
	altText             string
	postbackActionLabel string
	postbackActionData  string
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

	if !out.UserExists {
		presenter.replyMessage(msgFailAF, replyToken)
	} else if out.IsAlreadyAdded {
		presenter.replyMessage(msgAlreadyAddAF, replyToken)
	} else {
		presenter.replyMessage(msgSuccessAF, replyToken)
	}
}

// GetFavorites お気に入り一覧を送信する
func (presenter *LinePresenter) GetFavorites(out favoritedto.GetOutput) {
	msgs := carouselMsgs{
		noResult:            msgNoRegistGF,
		altText:             msgAltTextGF,
		postbackActionLabel: msgPostbackActionLabelGF,
		postbackActionData:  msgPostbackActionDataGF,
	}
	presenter.replyCarouselColumn(msgs, out.PlaceDetails, out.PlacePhotoURLs, out.ReplyToken)
}

// RemoveFavorite お気に入り削除結果を送信する
func (presenter *LinePresenter) RemoveFavorite(out favoritedto.RemoveOutput) {
	replyToken := out.ReplyToken

	if !out.UserExists {
		presenter.replyMessage(msgFailRF, replyToken)
	} else if out.IsAlreadyRemoved {
		presenter.replyMessage(msgAlreadyAddRF, replyToken)
	} else {
		presenter.replyMessage(msgSuccessRF, replyToken)
	}
}

// Search 検索結果を送信する
func (presenter *LinePresenter) Search(out searchdto.Output) {
	msgs := carouselMsgs{
		noResult:            msgNoRegistS,
		altText:             fmt.Sprintf(msgAltTextS, out.Q),
		postbackActionLabel: msgPostbackActionLabelS,
		postbackActionData:  msgPostbackActionDataS,
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

		data := fmt.Sprintf(msgs.postbackActionData, pd.PlaceID)
		cc := linebot.NewCarouselColumn(
			placePhotoURLs[i],
			pd.Name,
			formattedAddress,
			linebot.NewURIAction("Open Google Map", pd.URL),
			linebot.NewPostbackAction(msgs.postbackActionLabel, data, "", ""),
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
