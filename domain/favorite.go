package domain

import (
	"time"
)

// Favorite お気に入り
type Favorite struct {
	ID         uint      `gorm:"primary_key"`
	LineUserID uint      `json:"user_id"`
	PlaceID    string    `json:"video_id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`

	User User
}
