package message

import "github.com/pinfake/pes6go/data/block"

type RoomDeleted struct {
	Id uint32
}

func (m RoomDeleted) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4305, []block.Piece{
		block.Uint32{m.Id},
	})
}

func NewRoomDeleted(id uint32) RoomDeleted {
	return RoomDeleted{
		Id: id,
	}
}
