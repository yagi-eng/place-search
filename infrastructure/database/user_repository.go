package database

import (
	"github.com/yagi-eng/virtual-travel/domain/model"

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

// FindOne ユーザを検索する
func (repository *UserRepository) FindOne(lineUserID string) uint {
	user := model.User{}
	if repository.db.Table("users").
		Where(model.User{LineUserID: lineUserID}).First(&user).RecordNotFound() {

		return 0
	}
	return user.ID
}

// Save ユーザを登録する
func (repository *UserRepository) Save(lineUserID string) uint {
	user := model.User{}
	if repository.db.Table("users").
		Where(model.User{LineUserID: lineUserID}).First(&user).RecordNotFound() {

		user = model.User{LineUserID: lineUserID}
		repository.db.Create(&user)
	}
	return user.ID
}
