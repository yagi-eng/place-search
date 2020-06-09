package repository

// IUserRepository ユーザレポジトリインターフェース
type IUserRepository interface {
	Save(string) uint
	FindOne(string) uint
}
