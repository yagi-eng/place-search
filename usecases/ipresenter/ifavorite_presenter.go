package ipresenter

import "virtual-travel/usecases/dto/favoritedto"

// IFavoritePresenter お気に入りプレゼンター
type IFavoritePresenter interface {
	Add(favoritedto.AddOutput)
	Get(favoritedto.GetOutput)
}
