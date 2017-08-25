package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerParticipateResponse struct {
	*block.PlayerParticipateResponse
}

func (data PlayerParticipateResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4364, data.PlayerParticipateResponse)
}

func NewPlayerParticipateResponse(info *block.PlayerParticipateResponse) PlayerParticipateResponse {
	return PlayerParticipateResponse{info}
}
