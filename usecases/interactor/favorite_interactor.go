package interactor

import (
	"github.com/yagi-eng/place-search/domain/repository"
	"github.com/yagi-eng/place-search/usecases/dto/favoritedto"
	"github.com/yagi-eng/place-search/usecases/igateway"
	"github.com/yagi-eng/place-search/usecases/ipresenter"
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

// Get お気に入り全件を取得する
func (interactor *FavoriteInteractor) Get(in favoritedto.GetInput) {
	PlaceIDs := interactor.favoriteRepository.FindAll(in.LineUserID)

	googleMapOutputs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLs(PlaceIDs, true)

	out := favoritedto.GetOutput{
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.GetFavorites(out)
}

// Add お気に入りを追加する
func (interactor *FavoriteInteractor) Add(in favoritedto.AddInput) favoritedto.AddOutput {
	userID := interactor.userRepository.Save(in.LineUserID)

	var userExists bool
	var isAdded bool
	if userID == 0 {
		userExists = false
		isAdded = false
	} else {
		userExists = true
		isAdded = interactor.favoriteRepository.Save(userID, in.PlaceID)
	}

	out := favoritedto.AddOutput{
		ReplyToken:     in.ReplyToken,
		UserExists:     userExists,
		IsAlreadyAdded: !isAdded,
	}
	interactor.linePresenter.AddFavorite(out)

	return out
}

// Remove お気に入りを削除する
func (interactor *FavoriteInteractor) Remove(in favoritedto.RemoveInput) {
	userID := interactor.userRepository.FindOne(in.LineUserID)

	var userExists bool
	var isRemoved bool
	if userID == 0 {
		userExists = false
		isRemoved = false
	} else {
		userExists = true
		isRemoved = interactor.favoriteRepository.Delete(userID, in.PlaceID)
	}

	out := favoritedto.RemoveOutput{
		ReplyToken:       in.ReplyToken,
		UserExists:       userExists,
		IsAlreadyRemoved: !isRemoved,
	}
	interactor.linePresenter.RemoveFavorite(out)
}
