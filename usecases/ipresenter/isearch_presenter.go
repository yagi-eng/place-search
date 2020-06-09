package ipresenter

import (
	"virtual-travel/usecases/dto/searchdto"
)

// ISearchPresenter 検索プレゼンター
type ISearchPresenter interface {
	Hundle(searchdto.Output)
}
