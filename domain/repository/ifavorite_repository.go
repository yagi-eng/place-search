package repository

// IFavoriteRepository お気に入りレポジトリインターフェース
type IFavoriteRepository interface {
	Save(uint, string) bool
	FindAll(string) []string
	Delete(uint, string) bool
}
