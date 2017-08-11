package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type RoomPlayerLinks struct {
	RoomLinks []block.Piece
}

func (r RoomPlayerLinks) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4346, []block.Piece{block.Uint32{0}})...)
	blocks = append(blocks, block.GetBlocks(0x4347, r.RoomLinks)...)
	blocks = append(blocks, block.GetBlocks(0x4348, []block.Piece{block.Uint32{0}})...)
	return blocks
}

func NewRoomPlayerLinks(playerLinks []block.RoomPlayerLink) RoomPlayerLinks {
	return RoomPlayerLinks{
		RoomLinks: block.GetPieces(reflect.ValueOf(playerLinks)),
	}
}