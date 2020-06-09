package favoritedto

import "github.com/yagi-eng/virtual-travel/usecases/dto/googlemapdto"

// GetOutput DTO
type GetOutput struct {
	ReplyToken       string
	GoogleMapOutputs []googlemapdto.Output
}
