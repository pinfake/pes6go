package command

import (
	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/server"
)

type Login struct {
	Command
}

func (cmd Login) execute(c *client.Client) {
	authentication := block.Authentication{
		Key: server.Encrypt(cmd.data["key"].([]byte)),
	}
	c.WriteBlock(block.GetBlocks(0x3003, authentication)[0])
}
