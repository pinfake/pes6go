package message

import (
	"reflect"

	"fmt"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerSettings struct {
	PlayerId       uint32
	PlayerSettings []block.Piece
}

func (r PlayerSettings) GetBlocks() []block.Block {
	var blocks []block.Block
	// Primero va el player Id

	fmt.Printf("SETTTINGSDASDFASD!!!!: % x\n", r.PlayerSettings)

	blocks = append(blocks, block.GetBlocks(0x3087,
		[]block.Piece{block.PlayerSettingsHeader{r.PlayerId}})...,
	)
	blocks = append(blocks, block.GetBlocks(0x3088, r.PlayerSettings)...)
	blocks = append(blocks, block.NewBlock(0x3089, block.Zero{}))
	return blocks
}

func NewPlayerSettingsMessage(playerId uint32, playerSettings block.PlayerSettings) PlayerSettings {
	return PlayerSettings{
		PlayerId:       playerId,
		PlayerSettings: block.GetPieces(reflect.ValueOf(playerSettings)),
	}
}
