package discovery;

import "github.com/pinfake/pes6go/network/blocks"
import (
    "github.com/pinfake/pes6go/network/messages"
    "time"
)
type Response struct {
    Title string;
    Time time.Time;
    Text string;
    messages.Message
}

func (r Response) getBlocks() []blocks.Block {
    return []blocks.Block{
        blocks.NewBlock(0x0001, blocks.Zero{}),
        blocks.NewBlock(0x0002, blocks.Info{
            Time: r.Time, Title: r.Title, Text: r.Text,
        }),
        blocks.NewBlock(0x0003, blocks.Zero{}),
    }
}