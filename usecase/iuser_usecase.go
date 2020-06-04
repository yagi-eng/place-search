package usecase

import "virtual-travel/usecase/dto/userdto"

// IUserUseCase ユーザ登録ユースケース
type IUserUseCase interface {
	Create(userdto.UserCreateInput)
}
