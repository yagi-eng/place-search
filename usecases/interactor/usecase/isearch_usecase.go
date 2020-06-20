package usecase

import "github.com/yagi-eng/place-search/usecases/dto/searchdto"

// ISearchUseCase 検索ユースケース
type ISearchUseCase interface {
	Hundle(searchdto.Input) searchdto.Output
}
