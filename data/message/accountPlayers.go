package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type AccountPlayers struct {
	Profiles *block.AccountPlayers
}

func (data AccountPlayers) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3012, data.Profiles)
}

func NewAccountProfilesMessage(players *block.AccountPlayers) AccountPlayers {
	return AccountPlayers{players}
}
