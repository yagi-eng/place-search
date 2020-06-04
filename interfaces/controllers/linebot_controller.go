package controllers

import (
	"virtual-travel/infrastructure/database"
	"virtual-travel/interfaces/linebots"
	"virtual-travel/usecase"
	"virtual-travel/usecase/interactor"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	Interactor usecase.IUserUseCase
}

// NewLinebotController コンストラクタ
// TODO DIを導入して逆向きの依存を解消する
func NewLinebotController(db *gorm.DB) *LinebotController {
	return &LinebotController{
		Interactor: &interactor.UserInteractor{
			Repo: &database.UserRepository{
				DB: db,
			},
		},
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
					linebots.GetPlaceDetails(c, bot, event, message.Text)
				}
			} else if event.Type == linebot.EventTypePostback {
				linebots.AddFavorites(controller.Interactor, bot, event)
			}
		}

		return nil
	}
}
