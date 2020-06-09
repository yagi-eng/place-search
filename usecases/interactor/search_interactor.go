package interactor

import (
	"virtual-travel/usecases/dto/searchdto"
	"virtual-travel/usecases/igateway"
	"virtual-travel/usecases/ipresenter"
)

const spot = " 観光地"

// SearchInteractor 検索インタラクタ
type SearchInteractor struct {
	googleMapGateway igateway.IGoogleMapGateway
	searchPresenter  ipresenter.ISearchPresenter
}

// NewSearchInteractor コンストラクタ
func NewSearchInteractor(
	googleMapGateway igateway.IGoogleMapGateway,
	searchPresenter ipresenter.ISearchPresenter) *SearchInteractor {

	return &SearchInteractor{
		googleMapGateway: googleMapGateway,
		searchPresenter:  searchPresenter,
	}
}

// Hundle 検索する
func (interactor *SearchInteractor) Hundle(in searchdto.Input) {
	placeDetails, placePhotoURLs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsFromQuery(in.Q + spot)

	out := searchdto.Output{
		Q:              in.Q,
		ReplyToken:     in.ReplyToken,
		PlaceDetails:   placeDetails,
		PlacePhotoURLs: placePhotoURLs,
	}
	interactor.searchPresenter.Hundle(out)
}
