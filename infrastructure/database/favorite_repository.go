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
func (repository *FavoriteRepository) Save(userID uint, placeID string) bool {
	favorite := model.Favorite{}
	if repository.db.Table("favorites").
		Where(model.Favorite{UserID: userID, PlaceID: placeID}).First(&favorite).RecordNotFound() {

		favorite = model.Favorite{UserID: userID, PlaceID: placeID}
		repository.db.Create(&favorite)
		return false
	}

	return true
}

// FindAll お気に入り全件取得
func (repository *FavoriteRepository) FindAll(LineUserID string) []string {
	user := model.User{}
	repository.db.Table("users").Where(model.User{LineUserID: LineUserID}).First(&user)

	favorites := []model.Favorite{}
	repository.db.Model(&user).Related(&favorites)

	placeIDs := []string{}
	for _, favorite := range favorites {
		placeIDs = append(placeIDs, favorite.PlaceID)
	}

	return placeIDs
}
