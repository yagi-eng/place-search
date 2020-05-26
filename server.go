package main

import (
	"virtual-travel/infrastructure"
	"virtual-travel/infrastructure/middlewares"
	"virtual-travel/util/errmsg"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	errmsg.LogFatal(err)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.GoogleMapClient())

	// Routes
	infrastructure.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
