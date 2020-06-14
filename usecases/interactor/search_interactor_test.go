package interactor

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_gateway "github.com/yagi-eng/virtual-travel/mock/gateway"
	mock_ipresenter "github.com/yagi-eng/virtual-travel/mock/presenter"
	"github.com/yagi-eng/virtual-travel/usecases/dto/googlemapdto"
	"github.com/yagi-eng/virtual-travel/usecases/dto/searchdto"
)

// TestHundle1
// 正常系のみ
func TestHundle1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := "東京"
	gmo := googlemapdto.Output{}
	gmos := []googlemapdto.Output{gmo}

	mockGoogleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)
	mockGoogleMapGW.EXPECT().
		GetPlaceDetailsAndPhotoURLsFromQuery(q + spot).
		Return(gmos)

	expected := searchdto.Output{
		ReplyToken:       replyToken,
		Q:                q,
		GoogleMapOutputs: gmos,
	}
	mockLinePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	mockLinePrst.EXPECT().
		Search(expected).
		Return()

	interactor := NewSearchInteractor(mockGoogleMapGW, mockLinePrst)

	in := searchdto.Input{
		ReplyToken: replyToken,
		Q:          q,
	}
	interactor.Hundle(in)
}
