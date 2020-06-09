package usecase

import "virtual-travel/usecases/dto/favoritedto"

// IFavoriteUseCase お気に入りユースケース
type IFavoriteUseCase interface {
	Add(favoritedto.AddInput)
	Get(favoritedto.GetInput)
	Remove(favoritedto.RemoveInput)
}
