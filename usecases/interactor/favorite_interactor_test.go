package interactor

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yagi-eng/place-search/domain/model"
	mock_gateway "github.com/yagi-eng/place-search/mock/gateway"
	mock_ipresenter "github.com/yagi-eng/place-search/mock/presenter"
	mock_repository "github.com/yagi-eng/place-search/mock/repository"
	"github.com/yagi-eng/place-search/usecases/dto/favoritedto"
)

const userIDSuccess = uint(10)
const userIDFail = uint(0)
const lineUserID = "lineID123"
const placeID = "googlemapPlaceID123"
const replyToken = "hogeToken"
const userExists = true
const isAlreadyAdded = false
const isAlreadyRemoved = false

// TestGet1
// 正常系のみ
func TestGet1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	placeIDs := []string{placeID}
	gmo := model.Place{}
	gmos := []model.Place{gmo}

	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)

	mockFavoriteRepo := mock_repository.NewMockIFavoriteRepository(ctrl)
	mockFavoriteRepo.EXPECT().
		FindAll(lineUserID).
		Return(placeIDs)

	mockGoogleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)
	mockGoogleMapGW.EXPECT().
		GetPlaceDetailsAndPhotoURLs(placeIDs, true).
		Return(gmos)

	expected := favoritedto.GetOutput{
		ReplyToken:       replyToken,
		GoogleMapOutputs: gmos,
	}
	mockLinePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	mockLinePrst.EXPECT().
		GetFavorites(expected).
		Return()

	interactor := NewFavoriteInteractor(mockUserRepo, mockFavoriteRepo, mockGoogleMapGW, mockLinePrst)

	in := favoritedto.GetInput{
		ReplyToken: replyToken,
		LineUserID: lineUserID,
	}
	interactor.Get(in)
}

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

// TestRemove1
// ユーザが存在しており、お気に入り削除されていない場合
func TestRemove1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUserRepo.EXPECT().
		FindOne(lineUserID).
		Return(userIDSuccess)

	mockFavoriteRepo := mock_repository.NewMockIFavoriteRepository(ctrl)
	mockFavoriteRepo.EXPECT().
		Delete(userIDSuccess, placeID).
		Return(!isAlreadyRemoved)

	mockGoogleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)

	expected := favoritedto.RemoveOutput{
		ReplyToken:       replyToken,
		UserExists:       userExists,
		IsAlreadyRemoved: isAlreadyRemoved,
	}
	mockLinePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	mockLinePrst.EXPECT().
		RemoveFavorite(expected).
		Return()

	interactor := NewFavoriteInteractor(mockUserRepo, mockFavoriteRepo, mockGoogleMapGW, mockLinePrst)

	in := favoritedto.RemoveInput{
		ReplyToken: replyToken,
		LineUserID: lineUserID,
		PlaceID:    placeID,
	}
	interactor.Remove(in)
}

// TestRemove2
// ユーザ検索に失敗した場合
func TestRemove2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUserRepo.EXPECT().
		FindOne(lineUserID).
		Return(userIDFail)

	mockFavoriteRepo := mock_repository.NewMockIFavoriteRepository(ctrl)

	mockGoogleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)

	expected := favoritedto.RemoveOutput{
		ReplyToken:       replyToken,
		UserExists:       !userExists,
		IsAlreadyRemoved: !isAlreadyRemoved,
	}
	mockLinePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	mockLinePrst.EXPECT().
		RemoveFavorite(expected).
		Return()

	interactor := NewFavoriteInteractor(mockUserRepo, mockFavoriteRepo, mockGoogleMapGW, mockLinePrst)

	in := favoritedto.RemoveInput{
		ReplyToken: replyToken,
		LineUserID: lineUserID,
		PlaceID:    placeID,
	}
	interactor.Remove(in)
}
