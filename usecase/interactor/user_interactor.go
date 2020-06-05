package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase/dto/userdto"
)

// UserInteractor ユーザ登録インタラクタ
type UserInteractor struct {
	repo repository.IUserRepository
}

// NewUserInteractor コンストラクタ
func NewUserInteractor(repo repository.IUserRepository) *UserInteractor {
	return &UserInteractor{repo: repo}
}

// Create ユーザを登録する
func (ir *UserInteractor) Create(in userdto.UserCreateInput) {
	LineUserID := in.LineUserID
	ir.repo.Save(LineUserID)
}
