package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown308c struct {
}

func (Unknown308c) GetBlocks() []*block.Block {
	return block.GetBlocks(0x308d, block.Uint32{0})
}

func NewUnknown308c() Unknown308c {
	return Unknown308c{}
}
