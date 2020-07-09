package favoritedto

import "github.com/yagi-eng/place-search/domain/model"

// GetOutput DTO
type GetOutput struct {
	ReplyToken       string
	GoogleMapOutputs []model.Place
}
