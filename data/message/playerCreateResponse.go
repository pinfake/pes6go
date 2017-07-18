package message

import "github.com/pinfake/pes6go/data/block"

type PlayerCreateResponse struct {
	Code uint32
}

func (m PlayerCreateResponse) GetBlocks() []block.Block {
	return block.GetBlocks(0x3022, []block.Piece{
		block.Uint32{m.Code},
	})
}
