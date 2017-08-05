package message

import "github.com/pinfake/pes6go/data/block"

type LeaveLobby struct {
	Id uint32
}

func (m LeaveLobby) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4221, []block.Piece{
		block.Uint32{m.Id},
	})
}
