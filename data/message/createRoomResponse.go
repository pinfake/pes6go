package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type CreateRoomResponse struct {
}

func (CreateRoomResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4311, block.Uint32{0})
}

func NewCreateRoomResponse() CreateRoomResponse {
	return CreateRoomResponse{}
}
