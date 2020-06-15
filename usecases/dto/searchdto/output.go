package searchdto

import "github.com/yagi-eng/place-search/usecases/dto/googlemapdto"

// Output DTO
type Output struct {
	ReplyToken       string
	Q                string
	GoogleMapOutputs []googlemapdto.Output
}
