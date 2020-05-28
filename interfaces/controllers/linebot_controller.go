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
					replyPlaceURLByCarousel(c, bot, event, message.Text)
				}
			}
		}
		return nil
	}
}

func replyPlaceURLByCarousel(c echo.Context, bot *linebot.Client, event *linebot.Event, text string) {
	placeDetails := googlemap.GetPlaceDetails(c, text)

	if len(placeDetails) == 0 {
		res := linebot.NewTextMessage("検索結果は0件でした")
		if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
			logrus.Fatalf("Error LINEBOT replying message: %v", err)
		}
		return
	}

	ccs := []*linebot.CarouselColumn{}
	for _, pd := range placeDetails {
		cc := linebot.NewCarouselColumn(
			pd.Icon,
			pd.Name,
			pd.FormattedAddress,
			linebot.NewURIAction("Google Map", pd.URL),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}

	res := linebot.NewTemplateMessage(
		"「"+text+"」の検索結果です",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)
	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		logrus.Fatalf("Error LINEBOT replying message: %v", err)
	}
}
