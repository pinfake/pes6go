package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type AccountingInit struct {
}

func (r AccountingInit) GetBlocks() []block.Block {
	return []block.Block{
		block.NewBlock(0x3002, block.GenericBody{
			Data: []byte{
				0x38, 0x2b, 0x46, 0x47, 0x02, 0x4b, 0x2f, 0x68,
				0x56, 0x28, 0x3f, 0x53, 0x10, 0x87, 0x32, 0xa0,
			},
		}),
	}
}
