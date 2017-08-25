package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerSettings struct {
	PlayerId uint32
	*block.PlayerSettings
}

func (data PlayerSettings) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks,
		block.GetBlocks(0x3087, block.PlayerSettingsHeader{data.PlayerId})...,
	)
	blocks = append(blocks, block.GetBlocks(0x3088, data.PlayerSettings)...)
	blocks = append(blocks, block.GetBlocks(0x3089, block.Uint32{0})...)
	return blocks
}

func NewPlayerSettingsMessage(playerId uint32, info *block.PlayerSettings) PlayerSettings {
	return PlayerSettings{
		PlayerId:       playerId,
		PlayerSettings: info,
	}
}
