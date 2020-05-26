package infrastructure

import (
	"virtual-travel/interfaces/controllers"

	"github.com/labstack/echo"
)

// Init ルーティング設定
func Init(e *echo.Echo) {

	e.POST("/linebot/callback", controllers.ReplyByBot())

	// apiのメンテナンスは一旦中止、LINEBOT開発を優先
	g := e.Group("/api")
	{
		g.GET("/search", controllers.SearchLocations())
	}

}
