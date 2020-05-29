package main

import (
	"os"
	"virtual-travel/infrastructure"
	"virtual-travel/infrastructure/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	// Heroku上で動かすためエラー処理はしない
	godotenv.Load()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.GoogleMapClient())
	e.Use(middlewares.LineBotClient())

	// Routes
	infrastructure.Init(e)

	// Heroku用にPORT番号を環境変数から取得する
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
