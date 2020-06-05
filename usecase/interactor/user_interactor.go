package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase/dto/userdto"
)

// UserInteractor ユーザ登録インタラクタ
type UserInteractor struct {
	repository repository.IUserRepository
}

// NewUserInteractor コンストラクタ
func NewUserInteractor(repository repository.IUserRepository) *UserInteractor {
	return &UserInteractor{repository: repository}
}

// Create ユーザを登録する
func (interactor *UserInteractor) Create(in userdto.UserCreateInput) {
	LineUserID := in.LineUserID
	interactor.repository.Save(LineUserID)
}
