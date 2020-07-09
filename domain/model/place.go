package model

import (
	"time"
)

// Place プレイス
type Place struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `json:"name"`
	PlaceID   string    `json:"place_id"`
	Address   string    `json:"address"`
	URL       string    `json:"url"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Favorites []Favorite
}
