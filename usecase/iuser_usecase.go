package usecase

import "virtual-travel/usecase/dto/userdto"

// IUserUseCase ユーザユースケース
type IUserUseCase interface {
	Create(userdto.UserCreateInput) userdto.UserCreateOutput
}
