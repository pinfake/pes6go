package discovery

import (
	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
	"time"
)

const (
	responseQuery1 = 0x00002009 + iota
	responseQuery2
	responseQuery3
)

type Response struct {
	Title string
	Time  time.Time
	Text  string
	messages.Message
}

func (r Response) getBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(responseQuery1, blocks.Zero{}),
		blocks.NewBlock(responseQuery2, blocks.Info{
			Time: r.Time, Title: r.Title, Text: r.Text,
		}),
		blocks.NewBlock(responseQuery3, blocks.Zero{}),
	}
}
