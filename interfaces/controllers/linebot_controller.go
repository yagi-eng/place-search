package controllers

import (
	"strings"
	"virtual-travel/interfaces/controllers/linebots"
	"virtual-travel/usecase"
	"virtual-travel/usecase/dto/favoritedto"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	favoriteInteractor usecase.IFavoriteUseCase
}

// NewLinebotController コンストラクタ
func NewLinebotController(favoriteInteractor usecase.IFavoriteUseCase) *LinebotController {
	return &LinebotController{
		favoriteInteractor: favoriteInteractor,
	}
}

// CatchEvents LINEBOTに関する処理
func (controller *LinebotController) CatchEvents() echo.HandlerFunc {
	return func(c echo.Context) error {
		bot := c.Get("lbc").(*linebot.Client)

		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			logrus.Fatalf("Error LINEBOT parsing request: %v", err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					gmc := c.Get("gmc").(*maps.Client)
					msg := message.Text
					if msg == "お気に入り" {
						favoriteGetInput := favoritedto.GetInput{
							Gmc:        gmc,
							Bot:        bot,
							ReplyToken: event.ReplyToken,
							LineUserID: event.Source.UserID,
						}
						controller.favoriteInteractor.Get(favoriteGetInput)
					} else {
						linebots.GetPlaceDetails(gmc, bot, event, msg)
					}
				}
			} else if event.Type == linebot.EventTypePostback {
				dataMap := createDataMap(event.Postback.Data)

				if dataMap["action"] == "favorite" {
					favoriteAddInput := favoritedto.AddInput{
						Bot:        bot,
						ReplyToken: event.ReplyToken,
						LineUserID: event.Source.UserID,
						PlaceID:    dataMap["placeId"],
					}
					controller.favoriteInteractor.Add(favoriteAddInput)
				}
			}
		}

		return nil
	}
}

func createDataMap(q string) map[string]string {
	dataMap := make(map[string]string)

	dataArr := strings.Split(q, "&")
	for _, data := range dataArr {
		splitedData := strings.Split(data, "=")
		dataMap[splitedData[0]] = splitedData[1]
	}

	return dataMap
}
