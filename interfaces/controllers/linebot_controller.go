package controllers

import (
	"fmt"
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
					// replyLocationURL(c, bot, event, message.Text)
					fmt.Println(message.Text)
					testCarousel(c, bot, event)
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

func testCarousel(c echo.Context, bot *linebot.Client, event *linebot.Event) {
	resp := linebot.NewTemplateMessage(
		"this is a carousel template with imageAspectRatio,  imageSize and imageBackgroundColor",
		linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
				"this is menu",
				"description",
				linebot.NewURIAction("View detail", "http://example.com/page/111"),
			).WithImageOptions("#FFFFFF"),
			linebot.NewCarouselColumn(
				"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
				"this is menu",
				"description",
				linebot.NewURIAction("View detail", "http://example.com/page/111"),
			).WithImageOptions("#FFFFFF"),
		).WithImageOptions("rectangle", "cover"),
	)

	if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
