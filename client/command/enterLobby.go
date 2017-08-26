package command

import (
	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/data/block"
)

type EnterLobby struct {
	Command
}

func (cmd EnterLobby) execute(c *client.Client) {
	joinLobby := block.JoinLobby{
		LobbyId: cmd.data["lobbyId"].(byte),
		Ip1:     "127.0.0.1",
		Port1:   7777,
		Ip2:     "127.0.0.1",
		Port2:   8888,
	}
	c.WriteBlock(block.GetBlocks(0x4202, joinLobby)[0])
}
