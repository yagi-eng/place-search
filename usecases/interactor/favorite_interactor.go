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
	linePresenter      ipresenter.ILinePresenter
}

// NewFavoriteInteractor コンストラクタ
func NewFavoriteInteractor(
	userRepository repository.IUserRepository,
	favoriteRepository repository.IFavoriteRepository,
	googleMapGateway igateway.IGoogleMapGateway,
	linePresenter ipresenter.ILinePresenter) *FavoriteInteractor {

	return &FavoriteInteractor{
		userRepository:     userRepository,
		favoriteRepository: favoriteRepository,
		googleMapGateway:   googleMapGateway,
		linePresenter:      linePresenter,
	}
}

// Add お気に入りを追加する
func (interactor *FavoriteInteractor) Add(in favoritedto.AddInput) {
	userID := interactor.userRepository.Save(in.LineUserID)

	isSuccess := true
	if userID == 0 {
		isSuccess = false
		return
	}

	isAlreadyAdded := interactor.favoriteRepository.Save(userID, in.PlaceID)

	out := favoritedto.AddOutput{
		ReplyToken:     in.ReplyToken,
		IsSuccess:      isSuccess,
		IsAlreadyAdded: isAlreadyAdded,
	}
	interactor.linePresenter.AddFavorite(out)
}

// Get お気に入り全件を取得する
func (interactor *FavoriteInteractor) Get(in favoritedto.GetInput) {
	PlaceIDs := interactor.favoriteRepository.FindAll(in.LineUserID)

	placeDetails, placePhotoURLs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLs(PlaceIDs, true)

	out := favoritedto.GetOutput{
		ReplyToken:     in.ReplyToken,
		PlaceDetails:   placeDetails,
		PlacePhotoURLs: placePhotoURLs,
	}
	interactor.linePresenter.GetFavorites(out)
}

// Remove お気に入りを削除する
func (interactor *FavoriteInteractor) Remove(in favoritedto.RemoveInput) {
	userID := interactor.userRepository.FindOne(in.LineUserID)

	isSuccess := true
	if userID == 0 {
		isSuccess = false
		return
	}

	isAlreadyRemoved := interactor.favoriteRepository.Delete(userID, in.PlaceID)

	out := favoritedto.RemoveOutput{
		ReplyToken:       in.ReplyToken,
		IsSuccess:        isSuccess,
		IsAlreadyRemoved: isAlreadyRemoved,
	}
	interactor.linePresenter.RemoveFavorite(out)
}
