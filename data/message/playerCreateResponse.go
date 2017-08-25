package message

import "github.com/pinfake/pes6go/data/block"

type PlayerCreateResponse struct {
	Code uint32
}

func (data PlayerCreateResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3022, block.Uint32{data.Code})
}

func NewPlayerCreateResponse(code uint32) PlayerCreateResponse {
	return PlayerCreateResponse{code}
}
