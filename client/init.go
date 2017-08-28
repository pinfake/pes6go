package client

import (
	"github.com/pinfake/pes6go/data/block"
)

type Init struct {
	Command
}

func (cmd Init) Execute(c *Client) {
	c.WriteBlock(block.GetBlocks(0x3001, block.Void{})[0])
}
