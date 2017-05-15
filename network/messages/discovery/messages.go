package discovery

import (
	"time"

	"github.com/pinfake/pes6go/data"
	"github.com/pinfake/pes6go/network/blocks"
)

const (
	initResponseQuery1 = 0x00002009 + iota
	initResponseQuery2
	initResponseQuery3
)

const (
	serverTimeResponseQuery1 = 0x00002007
)

type ServerMessage struct {
	Message data.ServerMessage
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

func (r ServerMessage) GetBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(initResponseQuery1, blocks.Zero{}),
		r.Message.GetBlock(initResponseQuery2),
		blocks.NewBlock(initResponseQuery3, blocks.Zero{}),
	}
}
