package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type RoomsInLobby struct {
	Rooms []*block.Room
}

func (data RoomsInLobby) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4301, block.Uint32{0})...)
	blocks = append(blocks, block.GetBlocks(0x4302, data.Rooms)...)
	blocks = append(blocks, block.GetBlocks(0x4303, block.Uint32{0})...)

	return blocks
}

func NewRoomsInLobbyMessage(info []*block.Room) RoomsInLobby {
	return RoomsInLobby{info}
}
