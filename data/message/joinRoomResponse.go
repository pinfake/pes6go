package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type JoinRoomResponse struct {
	pieces []block.Piece
}

func (info JoinRoomResponse) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4321, info.pieces)...)

	return blocks
}

func NewJoinRoomResponse(code uint32, position byte) JoinRoomResponse {
	return JoinRoomResponse{
		pieces: block.GetPieces(reflect.ValueOf(
			block.JoinRoomResponse{
				Code:     code,
				Position: position,
			},
		)),
	}
}
