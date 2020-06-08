package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/interfaces/presenter"
	"virtual-travel/usecase/dto/favoritedto"
	"virtual-travel/usecase/igateway"
)

// FavoriteInteractor お気に入りインタラクタ
type FavoriteInteractor struct {
	userRepository     repository.IUserRepository
	favoriteRepository repository.IFavoriteRepository
	googleMapGateway   igateway.IGoogleMapGateway
}

// NewFavoriteInteractor コンストラクタ
func NewFavoriteInteractor(
	userRepository repository.IUserRepository,
	favoriteRepository repository.IFavoriteRepository,
	googleMapGateway igateway.IGoogleMapGateway) *FavoriteInteractor {

	return &FavoriteInteractor{
		userRepository:     userRepository,
		favoriteRepository: favoriteRepository,
		googleMapGateway:   googleMapGateway,
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
	presenter.Add(out)
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
	presenter.Get(out)
}
