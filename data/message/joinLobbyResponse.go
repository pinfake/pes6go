package message

import "github.com/pinfake/pes6go/data/block"

type JoinLobbyResponse struct {
	Code uint32
}

func (m JoinLobbyResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4203, []block.Piece{
		block.Uint32{m.Code},
	})
}
