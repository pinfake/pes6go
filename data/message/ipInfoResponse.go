package message

import "github.com/pinfake/pes6go/data/block"

type IpInfoResponse struct {
}

func (m IpInfoResponse) GetBlocks() []block.Block {
	return block.GetBlocks(0x4203, []block.Piece{
		block.Id{0},
	})
}
