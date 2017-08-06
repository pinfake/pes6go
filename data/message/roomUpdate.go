package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type RoomUpdate struct {
	Room []block.Piece
}

func (r RoomUpdate) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4306, r.Room)...)

	return blocks
}

func NewRoomUpdateMessage(room block.Room) RoomUpdate {
	return RoomUpdate{
		Room: block.GetPieces(reflect.ValueOf(room)),
	}
}
