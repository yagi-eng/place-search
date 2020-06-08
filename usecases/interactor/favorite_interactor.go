package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecases/dto/favoritedto"
	"virtual-travel/usecases/igateway"
	"virtual-travel/usecases/ipresenter"
)

// FavoriteInteractor お気に入りインタラクタ
type FavoriteInteractor struct {
	userRepository     repository.IUserRepository
	favoriteRepository repository.IFavoriteRepository
	googleMapGateway   igateway.IGoogleMapGateway
	favoritePresenter  ipresenter.IFavoritePresenter
}

// NewFavoriteInteractor コンストラクタ
func NewFavoriteInteractor(
	userRepository repository.IUserRepository,
	favoriteRepository repository.IFavoriteRepository,
	googleMapGateway igateway.IGoogleMapGateway,
	favoritePresenter ipresenter.IFavoritePresenter) *FavoriteInteractor {

	return &FavoriteInteractor{
		userRepository:     userRepository,
		favoriteRepository: favoriteRepository,
		googleMapGateway:   googleMapGateway,
		favoritePresenter:  favoritePresenter,
	}
}

// Add お気に入りを追加する
func (interactor *FavoriteInteractor) Add(in favoritedto.AddInput) {
	userID := interactor.userRepository.Save(in.LineUserID)

	isSuccess := true
	if userID == 0 {
		isSuccess = false
	}

	isAlreadyAdded := interactor.favoriteRepository.Save(userID, in.PlaceID)

	out := favoritedto.AddOutput{
		Bot:            in.Bot,
		ReplyToken:     in.ReplyToken,
		IsSuccess:      isSuccess,
		IsAlreadyAdded: isAlreadyAdded,
	}
	interactor.favoritePresenter.Add(out)
}

// Get お気に入り全件を取得する
func (interactor *FavoriteInteractor) Get(in favoritedto.GetInput) {
	PlaceIDs := interactor.favoriteRepository.FindAll(in.LineUserID)

	placeDetails, placePhotoURLs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLs(PlaceIDs, true)

	out := favoritedto.GetOutput{
		Bot:            in.Bot,
		ReplyToken:     in.ReplyToken,
		PlaceDetails:   placeDetails,
		PlacePhotoURLs: placePhotoURLs,
	}
	interactor.favoritePresenter.Get(out)
}
