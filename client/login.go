package client

import (
	"github.com/pinfake/pes6go/crypt"
	"github.com/pinfake/pes6go/data/block"
)

type Login struct {
	Command
}

func (cmd Login) Execute(c *Client) {
	authentication := block.Authentication{
		Key: crypt.Encrypt(crypt.PadWithZeros([]byte(cmd.Data["key"].(string)), 32)),
	}
	c.WriteBlock(block.GetBlocks(0x3003, authentication)[0])
}
