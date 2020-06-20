package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/yagi-eng/place-search/usecases/dto/favoritedto"
	"github.com/yagi-eng/place-search/usecases/dto/searchdto"
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

// Search クエリによる検索
func (controller *APIController) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("q")
		latStr := c.QueryParam("lat")
		lngStr := c.QueryParam("lng")
		addr := c.QueryParam("addr")

		lat, lng := float64(0), float64(0)
		if latStr != "" && lngStr != "" {
			var err error
			lat, err = strconv.ParseFloat(latStr, 64)
			lng, err = strconv.ParseFloat(lngStr, 64)

			if err != nil {
				logrus.Errorf("Error strconv: %v", err)
			}
		}

		in := searchdto.Input{
			Q:    q,
			Addr: addr,
			Lat:  lat,
			Lng:  lng,
		}
		out := controller.searchInteractor.Hundle(in)

		return c.JSON(fasthttp.StatusOK, out)
	}
}

// AddFavorites お気に入り追加
func (controller *APIController) AddFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO POSTで受け取る
		lineUserID := c.QueryParam("line_user_id")
		placeID := c.QueryParam("place_id")

		in := favoritedto.AddInput{
			LineUserID: lineUserID,
			PlaceID:    placeID,
		}
		out := controller.favoriteInteractor.Add(in)

		return c.JSON(fasthttp.StatusOK, out)
	}
}

// RemoveFavorites お気に入り削除
func (controller *APIController) RemoveFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO POSTで受け取る
		lineUserID := c.QueryParam("line_user_id")
		placeID := c.QueryParam("place_id")

		in := favoritedto.RemoveInput{
			LineUserID: lineUserID,
			PlaceID:    placeID,
		}
		out := controller.favoriteInteractor.Remove(in)

		return c.JSON(fasthttp.StatusOK, out)
	}
}
