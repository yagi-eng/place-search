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
func (interactor *SearchInteractor) Hundle(in searchdto.Input) searchdto.Output {
	outQ := ""
	var googleMapOutputs []googlemapdto.Output
	if isNomination(in.Q, in.Lat, in.Lng) {
		outQ = in.Q
		q := outQ + " " + os.Getenv("QUERY")
		googleMapOutputs = interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsWithQuery(q)
	} else if isOnlyLocaleInfo(in.Addr, in.Lat, in.Lng) {
		outQ = in.Addr
		q := os.Getenv("QUERY") + " " + outQ
		googleMapOutputs = interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsWithQueryLatLng(q, in.Lat, in.Lng)
	} else {
		logrus.Error("Error unexpected user request")
	}

	out := searchdto.Output{
		Q:                outQ,
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.Search(out)

	return out
}

func isNomination(q string, lat float64, lng float64) bool {
	return q != "" && lat == 0 && lng == 0
}

func isOnlyLocaleInfo(addr string, lat float64, lng float64) bool {
	return addr != "" && lat != 0 && lng != 0
}
