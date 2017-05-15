package discovery

import (
	"fmt"
	"time"

	"github.com/pinfake/pes6go/data"
	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
	"github.com/pinfake/pes6go/network/messages/common"
	"github.com/pinfake/pes6go/network/messages/discovery"
	"github.com/pinfake/pes6go/server"
)

var handleMap = map[uint16]func(blocks.Block, *server.Connection) messages.Message{
	0x2008: HandleDiscoveryInit,
	0x2006: HandleServerTime,
	0x2005: HandleQueryServers,
	0x2200: HandleRankUrlsQuery,
	0x0005: HandleKeepAlive,
	0x0003: HandleDisconnect,
}

type Server struct {
	server.Handler
}

type VoidMessage struct {
	messages.Message
}

func (m VoidMessage) getBlocks() []blocks.Block {
	return []blocks.Block{}
}

func HandleDiscoveryInit(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling discovery init")
	return discovery.ServerMessage{
		Message: data.ServerMessage{
			Time:  time.Date(2017, 1, 1, 12, 0, 0, 0, time.UTC),
			Title: "Hey, this is a title!",
			Text:  "Hey, this is the text, not so long!\nEsto es español amigo.",
		},
	}
}

func HandleQueryServers(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling query servers")
	return discovery.ServerListMessage{
		Servers: data.Servers{
			Servers: []data.Server{
				{7, "GROUP-SP/", "127.0.0.1", 10887, 0},
				{6, "SCHE-SP/", "127.0.0.1", 10887, 0},
				{4, "QUICK0-SP/", "127.0.0.1", 10887, 0},
				{4, "QUICK1-SP/", "127.0.0.1", 10887, 0},
				{8, "MENU03-SP/", "127.0.0.1", 12881, 0},
				{3, "TurboLobas Inc.", "127.0.0.1", 10900, 50},
				{2, "ACCT03-SP/", "127.0.0.1", 12881, 0},
				{1, "GATE-SP/", "127.0.0.1", 10887, 0},
			},
		},
	}
}

func HandleRankUrlsQuery(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling rank urls")
	return discovery.RankUrlListMessage{
		RankUrls: data.RankUrls{
			RankUrls: []data.RankUrl{
				{0x0000, "http://pes6web.winning-eleven.net/pes6e2/ranking/we10getrank.html"},
				{0x0100, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10getgrprank.html"},
				{0x0200, "http://pes6web.winning-eleven.net/pes6e2/ranking/we10RankingWeek.html"},
				{0x0300, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10GrpRankingWeek.html"},
				{0x0400, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10RankingCup.html"},
				{0x0500, "http://www.pes6j.net/server/we10getgrpboard.html"},
				{0x0600, "http://www.pes6j.net/server/we10getgrpinvitelist.html"},
			},
		},
	}
}

func HandleServerTime(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling server time")
	return discovery.ServerTime{
		Time: time.Now(),
	}
}

func HandleKeepAlive(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling a keep alive")
	return common.KeepAlive{}
}

func HandleDisconnect(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("Handling disconnect")
	return nil
}

func (s Server) HandleBlock(block blocks.Block, c *server.Connection) (messages.Message, error) {
	method, ok := handleMap[block.Header.Query]
	if !ok {
		return nil, fmt.Errorf("Unknown query!")
	}
	return method(block, c), nil
}

func Start() {
	fmt.Println("Here i am the discovery server!")
	s := Server{}
	server.Serve(10881, s)
}
