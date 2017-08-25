package message

import "github.com/pinfake/pes6go/data/block"

type LoginResponse struct {
	Code uint32
}

func (data LoginResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3004, block.Uint32{data.Code})
}

func NewLoginResponse(code uint32) LoginResponse {
	return LoginResponse{code}
}
