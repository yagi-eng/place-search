package interactor

import (
	"os"

	"github.com/yagi-eng/place-search/usecases/dto/searchdto"
	"github.com/yagi-eng/place-search/usecases/igateway"
	"github.com/yagi-eng/place-search/usecases/ipresenter"
)

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
	query := os.Getenv("QUERY")
	googleMapOutputs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsFromQuery(in.Q + " " + query)

	out := searchdto.Output{
		Q:                in.Q,
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.Search(out)
}
