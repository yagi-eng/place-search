package database

import (
	"github.com/yagi-eng/virtual-travel/domain/model"

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

// Save お気に入りを追加する
func (repository *FavoriteRepository) Save(userID uint, placeID string) bool {
	favorite := model.Favorite{}
	if repository.db.Table("favorites").
		Where(model.Favorite{UserID: userID, PlaceID: placeID}).First(&favorite).RecordNotFound() {

		favorite = model.Favorite{UserID: userID, PlaceID: placeID}
		repository.db.Create(&favorite)
		return true
	}

	return false
}

// Delete お気に入りを削除する
func (repository *FavoriteRepository) Delete(userID uint, placeID string) bool {
	favorite := model.Favorite{}
	if repository.db.Table("favorites").
		Where(model.Favorite{UserID: userID, PlaceID: placeID}).First(&favorite).RecordNotFound() {

		return false
	}

	repository.db.Delete(&favorite)
	return true
}

// FindAll お気に入りを全件取得する
func (repository *FavoriteRepository) FindAll(lineUserID string) []string {
	user := model.User{}
	repository.db.Table("users").Where(model.User{LineUserID: lineUserID}).First(&user)

	favorites := []model.Favorite{}
	repository.db.Model(&user).Related(&favorites)

	placeIDs := []string{}
	for _, favorite := range favorites {
		placeIDs = append(placeIDs, favorite.PlaceID)
	}

	return placeIDs
}
