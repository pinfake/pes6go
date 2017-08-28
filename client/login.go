package client

import (
	"github.com/pinfake/pes6go/data/block"
)

type Login struct {
	Command
}

func (cmd Login) Execute(c *Client) {
	authentication := block.Authentication{
		Key:      cmd.Data["key"].(string),
		Password: cmd.Data["password"].(string),
	}
	c.WriteBlock(block.GetBlocks(0x3003, authentication)[0])
}
