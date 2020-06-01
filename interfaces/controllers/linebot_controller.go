package controllers

import (
	"virtual-travel/interfaces/botreply"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// CatchEvents LINEBOTに関する処理
func CatchEvents() echo.HandlerFunc {
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
					botreply.GetPlaceDetails(c, bot, event, message.Text)
				}
			} else if event.Type == linebot.EventTypePostback {
				res := linebot.NewTextMessage("お気に入りに追加しました！（してない）")
				if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
					logrus.Fatalf("Error LINEBOT replying message: %v", err)
				}
			}
		}

		return nil
	}
}
