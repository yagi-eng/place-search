package database

import (
	"virtual-travel/domain/model"

	"github.com/jinzhu/gorm"
)

// userRepository ユーザレポジトリ
type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Save(LineUserID string) {
	repo.DB, _ = Connect()
	defer repo.DB.Close()

	// output sql query
	repo.DB.LogMode(true)

	user := model.User{}
	if repo.DB.Table("users").
		Where(model.User{LineUserID: LineUserID}).First(&user).RecordNotFound() {

		user = model.User{LineUserID: LineUserID}
		repo.DB.Create(&user)
	}
}
