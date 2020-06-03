package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase"
)

// UserCreateInteractor ユーザ登録インタラクター
type UserCreateInteractor struct {
	repo repository.IUserRepository
}

// Save ユーザを登録する
func (ir *UserCreateInteractor) Save(in usecase.UserCreateInput) {
	userLineID := in.LineUserID
	ir.repo.Save(userLineID)
}
