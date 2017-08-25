package message

import "github.com/pinfake/pes6go/data/block"

type PlayerIdResponse struct {
	PlayerIdResponse *block.PlayerIdResponse
}

func (data PlayerIdResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3062, data.PlayerIdResponse)
}

func NewPlayerIdResponseMessage() PlayerIdResponse {
	return PlayerIdResponse{&block.PlayerIdResponse{}}
}
