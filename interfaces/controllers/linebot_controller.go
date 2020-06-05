package controllers

import (
	"virtual-travel/interfaces/linebots"
	"virtual-travel/usecase"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	interactor usecase.IUserUseCase
}

// NewLinebotController コンストラクタ
func NewLinebotController(interactor usecase.IUserUseCase) *LinebotController {
	return &LinebotController{interactor: interactor}
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
					linebots.GetPlaceDetails(c, bot, event, message.Text)
				}
			} else if event.Type == linebot.EventTypePostback {
				linebots.AddFavorites(controller.interactor, bot, event)
			}
		}

		return nil
	}
}
