package discovery

import (
	"time"

	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
)

const (
	initResponseQuery1 = 0x00002009 + iota
	initResponseQuery2
	initResponseQuery3
)

const (
	serverTimeResponseQuery1 = 0x00002007
)

type Init struct {
	Title string
	Time  time.Time
	Text  string
	messages.Message
}

type ServerTime struct {
	Time time.Time
}

func (r ServerTime) GetBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(serverTimeResponseQuery1, blocks.ServerTime{
			Time: r.Time,
		}),
	}
}

func (r Init) GetBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(initResponseQuery1, blocks.Zero{}),
		blocks.NewBlock(initResponseQuery2, blocks.Info{
			Time: r.Time, Title: r.Title, Text: r.Text,
		}),
		blocks.NewBlock(initResponseQuery3, blocks.Zero{}),
	}
}
