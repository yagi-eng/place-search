package interactor

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yagi-eng/place-search/usecases/dto/googlemapdto"
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
	var googleMapOutputs []googlemapdto.Output
	if in.Q != "" {
		query := os.Getenv("QUERY")
		googleMapOutputs = interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsFromQuery(in.Q + " " + query)
	} else if in.Lat != 0 && in.Lng != 0 {
		googleMapOutputs = interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsFromLatLng(in.Lat, in.Lng)
	} else {
		logrus.Fatal("Error unexpected user request")
	}

	out := searchdto.Output{
		Q:                in.Q,
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.Search(out)
}
