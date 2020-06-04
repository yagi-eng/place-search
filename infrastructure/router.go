package infrastructure

import (
	"virtual-travel/infrastructure/database"
	"virtual-travel/interfaces/controllers"

	"github.com/labstack/echo"
)

// Init ルーティング設定
func Init(e *echo.Echo) {
	db, _ := database.Connect()
	defer db.Close()
	// output sql query
	db.LogMode(true)

	linebotController := controllers.NewLinebotController(db)

	e.POST("/linebot/callback", linebotController.CatchEvents())
}
