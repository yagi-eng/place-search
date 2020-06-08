package ipresenter

import "virtual-travel/usecase/dto/favoritedto"

// IFavoritePresenter お気に入りプレゼンター
type IFavoritePresenter interface {
	Add(favoritedto.AddOutput)
	Get(favoritedto.GetOutput)
}
