package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerJoinedLobby struct {
	Player []block.Piece
}

func (r PlayerJoinedLobby) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocksFromPieces(0x4220, r.Player)...)

	return blocks
}

func NewPlayerJoinedLobbyMessage(player block.Player) PlayerJoinedLobby {
	return PlayerJoinedLobby{
		Player: block.GetPieces(reflect.ValueOf(player)),
	}
}
