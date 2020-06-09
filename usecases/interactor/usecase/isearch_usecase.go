package usecase

import "github.com/yagi-eng/virtual-travel/usecases/dto/searchdto"

// ISearchUseCase 検索ユースケース
type ISearchUseCase interface {
	Hundle(searchdto.Input)
}
