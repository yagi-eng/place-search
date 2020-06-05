package database

import (
	"virtual-travel/domain/model"

	"github.com/jinzhu/gorm"
)

// FavoriteRepository お気に入りレポジトリ
type FavoriteRepository struct {
	db *gorm.DB
}

// NewFavoriteRepository コンストラクタ
func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

// Save お気に入り追加
func (repository *FavoriteRepository) Save(userID uint, placeID string) {
	favorite := model.Favorite{}
	if repository.db.Table("favorites").
		Where(model.Favorite{UserID: userID, PlaceID: placeID}).First(&favorite).RecordNotFound() {

		favorite = model.Favorite{UserID: userID, PlaceID: placeID}
		repository.db.Create(&favorite)
	}
}
