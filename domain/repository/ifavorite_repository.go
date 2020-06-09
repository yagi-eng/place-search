package repository

// IFavoriteRepository お気に入りレポジトリインターフェース
type IFavoriteRepository interface {
	FindAll(string) []string
	Save(uint, string) bool
	Delete(uint, string) bool
}
