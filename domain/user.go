package domain

import (
	"time"
)

// User ユーザ
type User struct {
	ID         uint       `gorm:"primary_key"`
	LineUserID string     `json:"-"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `sql:"index"json:"-"`

	Favorites []Favorite
}
