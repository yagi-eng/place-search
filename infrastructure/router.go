package infrastructure

import (
	"virtual-travel/interfaces/controllers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Init ルーティング設定
func Init(db *gorm.DB, e *echo.Echo) {
	linebotController := controllers.NewLinebotController(db)

	e.POST("/linebot/callback", linebotController.CatchEvents())
}
