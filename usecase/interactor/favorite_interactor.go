package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase/dto/favoritedto"
)

// FavoriteInteractor お気に入りインタラクタ
type FavoriteInteractor struct {
	repository repository.IFavoriteRepository
}

// NewFavoriteInteractor コンストラクタ
func NewFavoriteInteractor(repository repository.IFavoriteRepository) *FavoriteInteractor {
	return &FavoriteInteractor{repository: repository}
}

// Add お気に入りを追加する
func (interactor *FavoriteInteractor) Add(in favoritedto.FavoriteAddInput) {
	UserID := in.UserID
	PlaceID := in.PlaceID
	interactor.repository.Save(UserID, PlaceID)
}
