package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type RoomUpdate struct {
	*block.Room
}

func (data RoomUpdate) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4306, data.Room)
}

func NewRoomUpdateMessage(room *block.Room) RoomUpdate {
	return RoomUpdate{room}
}
