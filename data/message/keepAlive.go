package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type KeepAlive struct {
}

func (m KeepAlive) GetBlocks() []*block.Block {
	return block.GetBlocks(0x0005, []block.Piece{
		block.Void{},
	})
}
