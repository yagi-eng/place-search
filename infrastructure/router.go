package infrastructure

import (
	"github.com/yagi-eng/place-search/interfaces/controllers"

	"github.com/labstack/echo"
)

// Router ルーティング
type Router struct {
	e  *echo.Echo
	lc *controllers.LinebotController
	ac *controllers.APIController
}

// NewRouter コンストラクタ
func NewRouter(e *echo.Echo, lc *controllers.LinebotController, ac *controllers.APIController) *Router {
	return &Router{e: e, lc: lc, ac: ac}
}

// Init ルーティング設定
func (r *Router) Init() {
	r.e.POST("/linebot/callback", r.lc.CatchEvents())

	api := r.e.Group("/googlemap/api")
	{
		api.GET("/search", r.ac.Search())

		favorite := api.Group("/favorite")
		{
			favorite.GET("/add", r.ac.AddFavorites())
			favorite.GET("/remove", r.ac.RemoveFavorites())
		}
	}
}
