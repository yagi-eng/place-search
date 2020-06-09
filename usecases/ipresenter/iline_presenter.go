package ipresenter

import (
	"virtual-travel/usecases/dto/favoritedto"
	"virtual-travel/usecases/dto/searchdto"
)

// ILinePresenter LINEBOTプレゼンタ
type ILinePresenter interface {
	AddFavorite(favoritedto.AddOutput)
	GetFavorites(favoritedto.GetOutput)
	Search(searchdto.Output)
}
