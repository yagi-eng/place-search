package infrastructure

import (
	"virtual-travel/interfaces/controllers"

	"github.com/labstack/echo"
)

// Init ルーティング設定
func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.GET("/search", controllers.SearchLocations())
	}

}
