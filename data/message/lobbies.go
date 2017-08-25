package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Lobbies struct {
	*block.Lobbies
}

func (data Lobbies) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4201, data.Lobbies)
}

func NewLobbies(info *block.Lobbies) Lobbies {
	return Lobbies{info}
}
