package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerUpdate struct {
	*block.Player
}

func (data PlayerUpdate) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4222, data.Player)
}

func NewPlayerUpdate(player *block.Player) PlayerUpdate {
	return PlayerUpdate{player}
}
