package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type AccountProfiles struct {
	Profiles *block.AccountPlayers
}

func (data AccountProfiles) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3012, data.Profiles)
}

func NewAccountProfiles(players *block.AccountPlayers) AccountProfiles {
	return AccountProfiles{players}
}
