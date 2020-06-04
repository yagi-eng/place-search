package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase/dto/userdto"
)

// UserInteractor ユーザ登録インタラクタ
type UserInteractor struct {
	Repo repository.IUserRepository
}

// Create ユーザを登録する
func (ir *UserInteractor) Create(in userdto.UserCreateInput) {
	LineUserID := in.LineUserID
	ir.Repo.Save(LineUserID)
}
