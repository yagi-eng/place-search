package usecase

import "virtual-travel/usecases/dto/searchdto"

// ISearchUseCase 検索ユースケース
type ISearchUseCase interface {
	Hundle(searchdto.Input)
}
