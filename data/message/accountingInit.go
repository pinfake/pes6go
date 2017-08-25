package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type AccountingInit struct {
	*block.AccountingInit
}

func (data AccountingInit) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3002, data.AccountingInit)
}

func NewAccountingInit() AccountingInit {
	return AccountingInit{&block.AccountingInit{}}
}
