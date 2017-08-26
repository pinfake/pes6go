package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ReplayBlock struct {
	*block.Block
}

func (data ReplayBlock) GetBlocks() []*block.Block {
	return []*block.Block{data.Block}
}

func NewReplayBlockMessage(b *block.Block) ReplayBlock {
	return ReplayBlock{b}
}
