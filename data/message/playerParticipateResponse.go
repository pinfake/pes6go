package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerParticipateResponse struct {
	Response []block.Piece
}

func (r PlayerParticipateResponse) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4364, r.Response)...)

	return blocks
}

func NewPlayerParticipateResponse(info block.PlayerParticipateResponse) PlayerParticipateResponse {
	return PlayerParticipateResponse{
		Response: block.GetPieces(reflect.ValueOf(info)),
	}
}
