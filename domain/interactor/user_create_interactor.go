package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase"
)

// UserCreateInteractor ユーザ登録インタラクタ
type UserCreateInteractor struct {
	Repo repository.IUserRepository
}

// Handle ユーザを登録する
func (ir *UserCreateInteractor) Handle(in usecase.UserCreateInput) {
	LineUserID := in.LineUserID
	ir.Repo.Save(LineUserID)
}
