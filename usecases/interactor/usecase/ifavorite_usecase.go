package usecase

import "github.com/yagi-eng/place-search/usecases/dto/favoritedto"

// IFavoriteUseCase お気に入りユースケース
type IFavoriteUseCase interface {
	Get(favoritedto.GetInput) favoritedto.GetOutput
	Add(favoritedto.AddInput) favoritedto.AddOutput
	Remove(favoritedto.RemoveInput) favoritedto.RemoveOutput
}
