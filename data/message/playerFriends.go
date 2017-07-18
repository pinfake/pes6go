package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerFriends struct {
	PlayerFriends []block.Piece
}

func (r PlayerFriends) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.GetBlocks(0x3082, []block.Piece{block.Uint32{0}})...)
	blocks = append(blocks, block.NewBlock(0x3086, block.Void{}))
	return blocks
}

func NewPlayerFriendsMessage(playerFriends block.PlayerFriends) PlayerFriends {
	return PlayerFriends{
		PlayerFriends: block.GetPieces(reflect.ValueOf(playerFriends)),
	}
}
