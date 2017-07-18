package message

import "github.com/pinfake/pes6go/data/block"

type LoginResponse struct {
	Code uint32
}

func (m LoginResponse) GetBlocks() []block.Block {
	return block.GetBlocks(0x3004, []block.Piece{
		block.Uint32{m.Code},
	})
}
