package searchdto

import "virtual-travel/usecases/dto/googlemapdto"

// Output DTO
type Output struct {
	ReplyToken       string
	Q                string
	GoogleMapOutputs []googlemapdto.Output
}
