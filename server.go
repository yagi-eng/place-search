package main

import (
	"os"

	"github.com/yagi-eng/virtual-travel/infrastructure"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("ISPRD") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			logrus.Fatalf("Error loading env: %v", err)
		}
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// DB Connect
	db, err := infrastructure.Connect()
	if err != nil {
		logrus.Infof("Error connecting DB: %v", err)
		// Heroku用 アプリの起動に合わせてDBが起動できないことがあるので再接続を試みる
		db, _ = infrastructure.Connect()
	}

	defer db.Close()
	// output sql query
	db.LogMode(true)

	// Routes
	r := Initialize(e, db)
	r.Init()

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
