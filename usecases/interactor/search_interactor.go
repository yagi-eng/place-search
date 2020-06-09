package interactor

import (
	"github.com/yagi-eng/virtual-travel/usecases/dto/searchdto"
	"github.com/yagi-eng/virtual-travel/usecases/igateway"
	"github.com/yagi-eng/virtual-travel/usecases/ipresenter"
)

const spot = " 観光地"

// SearchInteractor 検索インタラクタ
type SearchInteractor struct {
	googleMapGateway igateway.IGoogleMapGateway
	linePresenter    ipresenter.ILinePresenter
}

// NewSearchInteractor コンストラクタ
func NewSearchInteractor(
	googleMapGateway igateway.IGoogleMapGateway,
	linePresenter ipresenter.ILinePresenter) *SearchInteractor {

	return &SearchInteractor{
		googleMapGateway: googleMapGateway,
		linePresenter:    linePresenter,
	}
}

// Hundle 検索する
func (interactor *SearchInteractor) Hundle(in searchdto.Input) {
	googleMapOutputs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsFromQuery(in.Q + spot)

	out := searchdto.Output{
		Q:                in.Q,
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.Search(out)
}
