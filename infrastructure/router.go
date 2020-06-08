package infrastructure

import (
	"virtual-travel/interfaces/controllers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Router ルーティング
type Router struct {
	e *echo.Echo
	// TODO いらない？
	db *gorm.DB
	c  *controllers.LinebotController
}

// NewRouter コンストラクタ
func NewRouter(e *echo.Echo, db *gorm.DB, c *controllers.LinebotController) *Router {
	return &Router{
		e:  e,
		db: db,
		c:  c,
	}
}

// Init ルーティング設定
func (r *Router) Init() {
	r.e.POST("/linebot/callback", r.c.CatchEvents())
}
