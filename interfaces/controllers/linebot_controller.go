package controllers

import (
	"virtual-travel/interfaces/googlemap"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// ReplyByBot LINEBOTに関する処理
func ReplyByBot() echo.HandlerFunc {
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
					replyLocationURL(c, bot, event, message.Text)
				}
			}
		}
		return nil
	}
}

func replyLocationURL(c echo.Context, bot *linebot.Client, event *linebot.Event, text string) {
	locationURLs := googlemap.GetLocationURLs(c, text)

	replyMsg := "検索結果は0件でした"
	if len(locationURLs) > 0 {
		replyMsg = locationURLs[0]
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
