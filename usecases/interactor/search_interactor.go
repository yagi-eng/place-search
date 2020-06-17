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
	if isNomination(in.Q, in.Lat, in.Lng) {
		q := in.Q + " " + os.Getenv("QUERY")
		googleMapOutputs = interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsWithQuery(q)
	} else if isLocalMessage(in.Q, in.Lat, in.Lng) {
		q := os.Getenv("QUERY") + " " + in.Q
		googleMapOutputs = interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsWithQueryLatLng(q, in.Lat, in.Lng)
	} else {
		logrus.Error("Error unexpected user request")
	}

	out := searchdto.Output{
		Q:                in.Q,
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.Search(out)
}

func isNomination(q string, lat float64, lng float64) bool {
	return q != "" && lat == 0 && lng == 0
}

func isLocalMessage(addr string, lat float64, lng float64) bool {
	return addr != "" && lat != 0 && lng != 0
}
