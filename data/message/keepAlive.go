package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type KeepAlive struct {
}

func (KeepAlive) GetBlocks() []*block.Block {
	return block.GetBlocks(0x0005, block.Void{})
}

func NewKeepAlive() KeepAlive {
	return KeepAlive{}
}
