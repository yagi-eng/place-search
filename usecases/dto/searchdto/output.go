package searchdto

import "github.com/yagi-eng/place-search/domain/model"

// Output DTO
type Output struct {
	ReplyToken       string
	Q                string
	GoogleMapOutputs []model.Place
}
