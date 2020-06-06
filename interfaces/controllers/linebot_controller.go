package controllers

import (
	"strings"
	"virtual-travel/interfaces/controllers/linebots"
	"virtual-travel/usecase"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	userinteractor     usecase.IUserUseCase
	favoriteinteractor usecase.IFavoriteUseCase
}

// NewLinebotController コンストラクタ
func NewLinebotController(userinteractor usecase.IUserUseCase, favoriteinteractor usecase.IFavoriteUseCase) *LinebotController {
	return &LinebotController{
		userinteractor:     userinteractor,
		favoriteinteractor: favoriteinteractor,
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
					gm := c.Get("gmc").(*maps.Client)
					linebots.GetPlaceDetails(gm, bot, event, message.Text)
				}
			} else if event.Type == linebot.EventTypePostback {
				dataMap := createDataMap(event.Postback.Data)

				if dataMap["action"] == "favorite" {
					placeID := dataMap["placeId"]
					linebots.AddFavorites(controller.userinteractor, controller.favoriteinteractor, bot, event, placeID)
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
