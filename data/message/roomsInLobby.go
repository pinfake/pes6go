package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type RoomsInLobby struct {
	Rooms []block.Piece
}

func (r RoomsInLobby) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocksFromPieces(0x4301, []block.Piece{
		block.Uint32{0},
	})...)
	blocks = append(blocks, block.GetBlocksFromPieces(0x4302, r.Rooms)...)
	blocks = append(blocks, block.GetBlocksFromPieces(0x4303, []block.Piece{
		block.Uint32{0},
	})...)

	return blocks
}

func NewRoomsInLobbyMessage(pieces []*block.Room) RoomsInLobby {
	return RoomsInLobby{
		Rooms: block.GetPieces(reflect.ValueOf(pieces)),
	}
}
