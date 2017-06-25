package message

import "github.com/pinfake/pes6go/data/block"

const loginResponseQuery = 0x3004

type LoginResponse struct {
}

func (m LoginResponse) GetBlocks() []block.Block {
	return []block.Block{
		block.NewBlock(loginResponseQuery, block.Zero{}),
	}
}
