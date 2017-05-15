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

type ServerListMessage struct {
	Servers data.Servers
}

type ServerTime struct {
	Time time.Time
}

type RankUrlListMessage struct {
	RankUrls data.RankUrls
}

func (r RankUrlListMessage) GetBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(0x2201, blocks.GenericBody{
			Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55},
		}),
		r.RankUrls.GetBlock(0x2202),
		blocks.NewBlock(0x2203, blocks.Zero{}),
	}
}

func (r ServerListMessage) GetBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(0x2002, blocks.Void{}),
		r.Servers.GetBlock(0x2003),
		blocks.NewBlock(0x2004, blocks.Void{}),
	}
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
