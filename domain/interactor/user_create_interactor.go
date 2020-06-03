package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase"
)

// UserCreateInteractor ユーザ登録インタラクタ
type userCreateInteractor struct {
	repo repository.IUserRepository
}

// Handle ユーザを登録する
func (ir *userCreateInteractor) Handle(in usecase.UserCreateInput) {
	LineUserID := in.LineUserID
	ir.repo.Save(LineUserID)
}
