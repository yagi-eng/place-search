package favoritedto

import "virtual-travel/usecases/dto/googlemapdto"

// GetOutput DTO
type GetOutput struct {
	ReplyToken       string
	GoogleMapOutputs []googlemapdto.Output
}
