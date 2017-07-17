package message

import "github.com/pinfake/pes6go/data/block"

const Ok = 0
const VerificationError = 0xffffff10
const ServiceUnavailableError = 0xffffff12

type LoginResponse struct {
	RCode uint32
}

func (m LoginResponse) GetBlocks() []block.Block {
	return block.GetBlocks(0x3004, []block.Piece{
		block.Id{m.RCode},
	})
}
