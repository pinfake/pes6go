package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type CreateRoomResponse struct {
}

func (r CreateRoomResponse) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4311, []block.Piece{block.Uint32{0}})...)
	return blocks
}

func NewCreateRoomResponse() CreateRoomResponse {
	return CreateRoomResponse{}
}
