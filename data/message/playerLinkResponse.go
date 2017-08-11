package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerLinkResponse struct {
	PlayerLink []block.Piece
}

func (r PlayerLinkResponse) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4b01, r.PlayerLink)...)

	return blocks
}

func NewPlayerLinkResponse(info block.PlayerLink) PlayerLinkResponse {
	return PlayerLinkResponse{
		PlayerLink: block.GetPieces(reflect.ValueOf(info)),
	}
}
