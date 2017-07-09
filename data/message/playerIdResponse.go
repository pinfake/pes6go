package message

import "github.com/pinfake/pes6go/data/block"

type PlayerIdResponse struct {
	PlayerIdResponse []block.Piece
}

func (m PlayerIdResponse) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.GetBlocks(0x3062,
		m.PlayerIdResponse)...)
	return blocks
}

func NewPlayerIdResponseMessage(code uint16) PlayerIdResponse {
	return PlayerIdResponse{
		PlayerIdResponse: []block.Piece{
			block.PlayerIdResponse{
				Code: code,
			},
		},
	}
}
