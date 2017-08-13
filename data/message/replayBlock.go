package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ReplayBlock struct {
	b *block.Block
}

func (r ReplayBlock) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, r.b)

	return blocks
}

func NewReplayBlockMessage(b *block.Block) ReplayBlock {
	return ReplayBlock{
		b: b,
	}
}
