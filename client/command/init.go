package command

import (
	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/data/block"
)

type Init struct {
	Command
}

func (cmd Init) execute(c *client.Client) {
	c.WriteBlock(block.GetBlocks(0x3001, block.Void{})[0])
}
