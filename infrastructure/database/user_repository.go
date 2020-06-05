package database

import (
	"virtual-travel/domain/model"

	"github.com/jinzhu/gorm"
)

// UserRepository ユーザレポジトリ
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository コンストラクタ
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save 登録
func (repository *UserRepository) Save(LineUserID string) uint {
	user := model.User{}
	if repository.db.Table("users").
		Where(model.User{LineUserID: LineUserID}).First(&user).RecordNotFound() {

		user = model.User{LineUserID: LineUserID}
		repository.db.Create(&user)
	}

	return user.ID
}
