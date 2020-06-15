package favoritedto

import "github.com/yagi-eng/place-search/usecases/dto/googlemapdto"

// GetOutput DTO
type GetOutput struct {
	ReplyToken       string
	GoogleMapOutputs []googlemapdto.Output
}
