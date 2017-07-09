package message

import "github.com/pinfake/pes6go/data/block"

type AccountPlayers struct {
	AccountPlayers []block.Piece
}

func (r AccountPlayers) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.GetBlocks(0x3012, r.AccountPlayers)...)

	return blocks
}

func NewAccountPlayersMessage(players block.AccountPlayers) AccountPlayers {
	return AccountPlayers{
		AccountPlayers: []block.Piece{
			players,
		},
	}
}
