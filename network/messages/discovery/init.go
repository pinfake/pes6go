package discovery;

import "github.com/pinfake/pes6go/network/blocks"
import (
    "github.com/pinfake/pes6go/network/messages"
    "time"
)
type Response struct {
    title string;
    time time.Time;
    text string;
    messages.Message
}

func (r Response) getBlocks() []blocks.Block {
    return []blocks.Block{
        blocks.NewBlock(0x0001, blocks.Zero{}),
        blocks.NewBlock(0x0002, blocks.Info{
            r.time, r.title, r.text,
        }),
        blocks.NewBlock(0x0003, blocks.Zero{}),
    }
}