package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type CreateRoomResponse struct {
	RoomLinks []block.Piece
}

func (r CreateRoomResponse) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4311, []block.Piece{block.Uint32{0}})...)
	blocks = append(blocks, block.GetBlocks(0x4346, []block.Piece{block.Uint32{0}})...)
	blocks = append(blocks, block.GetBlocks(0x4347, r.RoomLinks)...)
	blocks = append(blocks, block.GetBlocks(0x4348, []block.Piece{block.Uint32{0}})...)
	return blocks
}

func NewCreateRoomResponse(playerId uint32, playerSettings block.PlayerSettings) PlayerSettings {
	return PlayerSettings{
		PlayerId:       playerId,
		PlayerSettings: block.GetPieces(reflect.ValueOf(playerSettings)),
	}
}
