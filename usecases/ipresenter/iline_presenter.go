package ipresenter

import (
	"github.com/yagi-eng/virtual-travel/usecases/dto/favoritedto"
	"github.com/yagi-eng/virtual-travel/usecases/dto/searchdto"
)

// ILinePresenter LINEBOTプレゼンタ
type ILinePresenter interface {
	AddFavorite(favoritedto.AddOutput)
	GetFavorites(favoritedto.GetOutput)
	RemoveFavorite(favoritedto.RemoveOutput)
	Search(searchdto.Output)
}
