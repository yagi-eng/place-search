package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase/dto/userdto"
)

// UserInteractor ユーザインタラクタ
type UserInteractor struct {
	repository repository.IUserRepository
}

// NewUserInteractor コンストラクタ
func NewUserInteractor(repository repository.IUserRepository) *UserInteractor {
	return &UserInteractor{repository: repository}
}

// Create ユーザを登録する
func (interactor *UserInteractor) Create(in userdto.UserCreateInput) userdto.UserCreateOutput {
	lineUserID := in.LineUserID
	userID := interactor.repository.Save(lineUserID)

	return userdto.UserCreateOutput{UserID: userID}
}
