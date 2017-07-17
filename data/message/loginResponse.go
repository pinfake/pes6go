package message

import "github.com/pinfake/pes6go/data/block"

const LoginOk = 0
const LoginFail = 0xffffff10

type LoginResponse struct {
	RCode uint32
}

func (m LoginResponse) GetBlocks() []block.Block {
	return block.GetBlocks(0x3004, []block.Piece{
		block.Id{m.RCode},
	})
}
