package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerLinkResponse struct {
	*block.PlayerLink
}

func (data PlayerLinkResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4b01, data.PlayerLink)
}

func NewPlayerLinkResponse(info *block.PlayerLink) PlayerLinkResponse {
	return PlayerLinkResponse{info}
}
