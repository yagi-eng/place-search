package ipresenter

import (
	"virtual-travel/usecases/dto/favoritedto"
	"virtual-travel/usecases/dto/searchdto"
)

// ILinePresenter LINEBOTプレゼンタ
type ILinePresenter interface {
	AddFavorite(favoritedto.AddOutput)
	GetFavorite(favoritedto.GetOutput)
	Search(searchdto.Output)
}
