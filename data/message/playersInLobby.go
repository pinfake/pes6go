package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayersInLobby struct {
	Players []block.Piece
}

func (r PlayersInLobby) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4211, []block.Piece{
		block.Uint32{0},
	})...)
	blocks = append(blocks, block.GetBlocks(0x4212, r.Players)...)
	blocks = append(blocks, block.GetBlocks(0x4213, []block.Piece{
		block.Uint32{0},
	})...)

	return blocks
}

func NewPlayersInLobbyMessage(players []*block.Player) PlayersInLobby {
	return PlayersInLobby{
		Players: block.GetPieces(reflect.ValueOf(players)),
	}
}
