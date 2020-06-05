package repository

// IFavoriteRepository お気に入りレポジトリインターフェース
type IFavoriteRepository interface {
	Save(uint, string)
}
