package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerJoinedLobby struct {
	*block.Player
}

func (data PlayerJoinedLobby) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4220, data.Player)
}

func NewPlayerJoinedLobby(info *block.Player) PlayerJoinedLobby {
	return PlayerJoinedLobby{info}
}
