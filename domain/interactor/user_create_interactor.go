package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase"
)

// UserCreateInteractor ユーザ登録インタラクター
type userCreateInteractor struct {
	repo repository.IUserRepository
}

// Handle ユーザを登録する
func (ir *userCreateInteractor) Handle(in usecase.UserCreateInput) {
	userLineID := in.LineUserID
	ir.repo.Save(userLineID)
}
