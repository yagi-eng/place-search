package infrastructure

import (
	"virtual-travel/interfaces/controllers"

	"github.com/labstack/echo"
)

// Router ルーティング
type Router struct {
	e *echo.Echo
	c *controllers.LinebotController
}

// NewRouter コンストラクタ
func NewRouter(e *echo.Echo, c *controllers.LinebotController) *Router {
	return &Router{e: e, c: c}
}

// Init ルーティング設定
func (r *Router) Init() {
	r.e.POST("/linebot/callback", r.c.CatchEvents())
}
