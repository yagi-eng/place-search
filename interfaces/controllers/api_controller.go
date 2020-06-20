package controllers

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/yagi-eng/place-search/usecases/interactor/usecase"

	"github.com/line/line-bot-sdk-go/linebot"
)

// APIController APIコントローラ
type APIController struct {
	favoriteInteractor usecase.IFavoriteUseCase
	searchInteractor   usecase.ISearchUseCase
	bot                *linebot.Client
}

// NewAPIController コンストラクタ
func NewAPIController(
	favoriteInteractor usecase.IFavoriteUseCase,
	searchInteractor usecase.ISearchUseCase) *APIController {

	return &APIController{
		favoriteInteractor: favoriteInteractor,
		searchInteractor:   searchInteractor,
	}
}

// SearchWithQuery クエリによる検索
func (controller *APIController) SearchWithQuery() echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("q")
		return c.JSON(fasthttp.StatusOK, q)
	}
}
