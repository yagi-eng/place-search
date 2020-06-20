//+build wireinject

package main

import (
	"github.com/yagi-eng/place-search/domain/repository"
	"github.com/yagi-eng/place-search/infrastructure"
	"github.com/yagi-eng/place-search/infrastructure/database"
	"github.com/yagi-eng/place-search/interfaces/controllers"
	"github.com/yagi-eng/place-search/interfaces/gateway"
	"github.com/yagi-eng/place-search/interfaces/presenter"
	"github.com/yagi-eng/place-search/usecases/igateway"
	"github.com/yagi-eng/place-search/usecases/interactor"
	"github.com/yagi-eng/place-search/usecases/interactor/usecase"
	"github.com/yagi-eng/place-search/usecases/ipresenter"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var superSet = wire.NewSet(
	// Database
	database.NewFavoriteRepository,
	wire.Bind(new(repository.IFavoriteRepository), new(*database.FavoriteRepository)),
	database.NewUserRepository,
	wire.Bind(new(repository.IUserRepository), new(*database.UserRepository)),

	// Gateway
	gateway.NewGoogleMapGateway,
	wire.Bind(new(igateway.IGoogleMapGateway), new(*gateway.GoogleMapGateway)),

	// Presenter
	presenter.NewLinePresenter,
	wire.Bind(new(ipresenter.ILinePresenter), new(*presenter.LinePresenter)),

	// Interactor
	interactor.NewFavoriteInteractor,
	wire.Bind(new(usecase.IFavoriteUseCase), new(*interactor.FavoriteInteractor)),
	interactor.NewSearchInteractor,
	wire.Bind(new(usecase.ISearchUseCase), new(*interactor.SearchInteractor)),

	// Controller
	controllers.NewLinebotController,
	controllers.NewAPIController,

	// Router
	infrastructure.NewRouter,
)

// Initialize DI
func Initialize(e *echo.Echo, db *gorm.DB) *infrastructure.Router {
	wire.Build(superSet)
	return &infrastructure.Router{}
}
