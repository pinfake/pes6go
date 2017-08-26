package message

import "github.com/pinfake/pes6go/data/block"

type RoomDeleted struct {
	Id uint32
}

func (data RoomDeleted) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4305, block.Uint32{data.Id})
}

func NewRoomDeleted(id uint32) RoomDeleted {
	return RoomDeleted{id}
}
