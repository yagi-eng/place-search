package usecase

import "github.com/yagi-eng/virtual-travel/usecases/dto/favoritedto"

// IFavoriteUseCase お気に入りユースケース
type IFavoriteUseCase interface {
	Get(favoritedto.GetInput)
	Add(favoritedto.AddInput)
	Remove(favoritedto.RemoveInput)
}
