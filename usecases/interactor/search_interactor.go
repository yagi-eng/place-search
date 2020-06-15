package interactor

import (
	"os"

	"github.com/sirupsen/logrus"
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
	q := ""
	if in.Q != "" {
		q = in.Q + " " + os.Getenv("QUERY")
	} else if in.Addr != "" {
		q = os.Getenv("QUERY") + " " + in.Addr
	} else {
		logrus.Fatal("Error unexpected user request")
	}

	googleMapOutputs := interactor.googleMapGateway.GetPlaceDetailsAndPhotoURLsFromQuery(q)

	out := searchdto.Output{
		Q:                in.Q,
		ReplyToken:       in.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	interactor.linePresenter.Search(out)
}
