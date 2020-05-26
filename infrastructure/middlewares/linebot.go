package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// LineBotClient LINEBOTインスタンスを生成
func LineBotClient() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			secret := os.Getenv("LBOT_SECRET")
			token := os.Getenv("LBOT_TOKEN")

			bot, err := linebot.New(secret, token)
			if err != nil {
				logrus.Fatalf("Error creating LINEBOT client: %v", err)
			}

			c.Set("lbc", bot)

			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}
