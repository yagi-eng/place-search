package controllers

import (
	"os"
	"strings"
	"virtual-travel/usecases/dto/favoritedto"
	"virtual-travel/usecases/interactor/usecase"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	favoriteInteractor usecase.IFavoriteUseCase
	bot                *linebot.Client
}

// NewLinebotController コンストラクタ
func NewLinebotController(favoriteInteractor usecase.IFavoriteUseCase) *LinebotController {
	secret := os.Getenv("LBOT_SECRET")
	token := os.Getenv("LBOT_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client: %v", err)
	}

	return &LinebotController{
		favoriteInteractor: favoriteInteractor,
		bot:                bot,
	}
}

// CatchEvents LINEBOTに関する処理
func (controller *LinebotController) CatchEvents() echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := controller.bot.ParseRequest(c.Request())
		if err != nil {
			logrus.Fatalf("Error LINEBOT parsing request: %v", err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					msg := message.Text
					if msg == "お気に入り" {
						favoriteGetInput := favoritedto.GetInput{
							ReplyToken: event.ReplyToken,
							LineUserID: event.Source.UserID,
						}
						controller.favoriteInteractor.Get(favoriteGetInput)
					} else {
						// linebots.GetPlaceDetails(gmc, bot, event, msg)
					}
				}
			} else if event.Type == linebot.EventTypePostback {
				dataMap := createDataMap(event.Postback.Data)

				if dataMap["action"] == "favorite" {
					favoriteAddInput := favoritedto.AddInput{
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
