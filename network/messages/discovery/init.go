package discovery;

import "github.com/pinfake/pes6go/network/blocks"
import "github.com/pinfake/pes6go/network/messages"
type Response struct {
    messages.Message
}

func (Response) getBlocks() []blocks.Block {
    return []blocks.Block{
        blocks.NewBlock(0x0001, blocks.Zero{}),
    }
}