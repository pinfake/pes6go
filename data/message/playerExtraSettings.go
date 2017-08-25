package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerExtraSettings struct {
	*block.PlayerExtraSettings
}

func (data PlayerExtraSettings) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4101, data.PlayerExtraSettings)
}

func NewPlayerExtraSettingsMessage(info *block.PlayerExtraSettings) PlayerExtraSettings {
	return PlayerExtraSettings{info}
}
