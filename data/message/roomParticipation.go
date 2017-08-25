package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type RoomParticipation struct {
	RoomParticipation []block.Piece
}

func (r RoomParticipation) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocksFromPieces(0x4365, r.RoomParticipation)...)

	return blocks
}

func NewRoomParticipation(info block.RoomParticipation) RoomParticipation {
	return RoomParticipation{
		RoomParticipation: block.GetPieces(reflect.ValueOf(info)),
	}
}
