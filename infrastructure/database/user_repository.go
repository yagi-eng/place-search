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

// Save ユーザ登録
func (repo *UserRepository) Save(LineUserID string) {
	user := model.User{}
	if repo.db.Table("users").
		Where(model.User{LineUserID: LineUserID}).First(&user).RecordNotFound() {

		user = model.User{LineUserID: LineUserID}
		repo.db.Create(&user)
	}
}
