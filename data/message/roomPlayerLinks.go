package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type RoomPlayerLinks struct {
	*block.RoomPlayerLinks
}

func (data RoomPlayerLinks) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4346, block.Uint32{0})...)
	blocks = append(blocks, block.GetBlocks(0x4347, data.RoomPlayerLinks)...)
	blocks = append(blocks, block.GetBlocks(0x4348, block.Uint32{0})...)
	return blocks
}

func NewRoomPlayerLinks(info *block.RoomPlayerLinks) RoomPlayerLinks {
	return RoomPlayerLinks{info}
}
