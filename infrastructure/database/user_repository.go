package database

import (
	"virtual-travel/domain/model"

	"github.com/jinzhu/gorm"
)

// UserRepository ユーザレポジトリ
type UserRepository struct {
	DB *gorm.DB
}

// Save ユーザ登録
func (repo *UserRepository) Save(LineUserID string) {
	user := model.User{}
	if repo.DB.Table("users").
		Where(model.User{LineUserID: LineUserID}).First(&user).RecordNotFound() {

		user = model.User{LineUserID: LineUserID}
		repo.DB.Create(&user)
	}
}
