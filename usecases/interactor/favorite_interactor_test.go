package interactor

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_gateway "github.com/yagi-eng/virtual-travel/mock/gateway"
	mock_ipresenter "github.com/yagi-eng/virtual-travel/mock/presenter"
	mock_repository "github.com/yagi-eng/virtual-travel/mock/repository"
	"github.com/yagi-eng/virtual-travel/usecases/dto/favoritedto"
)

const userIDSuccess = uint(10)
const userIDFail = uint(0)
const lineUserID = "lineID123"
const placeID = "googlemapPlaceID123"
const replyToken = "hogeToken"
const userExists = true
const isAlreadyAdded = false

// TestAdd1
// ユーザは存在しておらず、お気に入り登録されていない場合
func TestAdd1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUserRepo.EXPECT().
		Save(lineUserID).
		Return(userIDSuccess)

	mockFavoriteRepo := mock_repository.NewMockIFavoriteRepository(ctrl)
	mockFavoriteRepo.EXPECT().
		Save(userIDSuccess, placeID).
		Return(!isAlreadyAdded)

	mockGoogleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)

	expected := favoritedto.AddOutput{
		ReplyToken:     replyToken,
		UserExists:     userExists,
		IsAlreadyAdded: isAlreadyAdded,
	}
	mockLinePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	mockLinePrst.EXPECT().
		AddFavorite(expected).
		Return()

	interactor := NewFavoriteInteractor(mockUserRepo, mockFavoriteRepo, mockGoogleMapGW, mockLinePrst)

	in := favoritedto.AddInput{
		ReplyToken: replyToken,
		LineUserID: lineUserID,
		PlaceID:    placeID,
	}
	interactor.Add(in)
}

// TestAdd2
// ユーザは存在しており、お気に入り登録されている場合
func TestAdd2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUserRepo.EXPECT().
		Save(lineUserID).
		Return(userIDSuccess)

	mockFavoriteRepo := mock_repository.NewMockIFavoriteRepository(ctrl)
	mockFavoriteRepo.EXPECT().
		Save(userIDSuccess, placeID).
		Return(isAlreadyAdded)

	googleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)

	expected := favoritedto.AddOutput{
		ReplyToken:     replyToken,
		UserExists:     userExists,
		IsAlreadyAdded: !isAlreadyAdded,
	}
	linePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	linePrst.EXPECT().
		AddFavorite(expected).
		Return()

	interactor := NewFavoriteInteractor(mockUserRepo, mockFavoriteRepo, googleMapGW, linePrst)

	in := favoritedto.AddInput{
		ReplyToken: replyToken,
		LineUserID: lineUserID,
		PlaceID:    placeID,
	}
	interactor.Add(in)
}

// TestAdd3
// ユーザ登録に失敗した場合
func TestAdd3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUserRepo.EXPECT().
		Save(lineUserID).
		Return(userIDFail)

	mockFavoriteRepo := mock_repository.NewMockIFavoriteRepository(ctrl)
	googleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)

	expected := favoritedto.AddOutput{
		ReplyToken:     replyToken,
		UserExists:     !userExists,
		IsAlreadyAdded: !isAlreadyAdded,
	}
	linePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	linePrst.EXPECT().
		AddFavorite(expected).
		Return()

	interactor := NewFavoriteInteractor(mockUserRepo, mockFavoriteRepo, googleMapGW, linePrst)

	in := favoritedto.AddInput{
		ReplyToken: replyToken,
		LineUserID: lineUserID,
		PlaceID:    placeID,
	}
	interactor.Add(in)
}
