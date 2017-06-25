package discovery

import (
	"fmt"
	"time"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
)

var handleMap = map[uint16]func(block.Block, *server.Connection) message.Message{
	0x2008: HandleDiscoveryInit,
	0x2006: HandleServerTime,
	0x2005: HandleQueryServers,
	0x2200: HandleRankUrlsQuery,
	0x0005: HandleKeepAlive,
	0x0003: HandleDisconnect,
}

func HandleDiscoveryInit(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling discovery init")
	return message.Motd{
		Messages: []block.Piece{
			block.ServerMessage{
				Time:  time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
				Title: "Mariano Powered:",
				Text: "Es el vecino el que elige al alcalde y es el alcalde el que quiere " +
					"que sean los vecinos el alcalde",
			},
		},
	}
}

func HandleQueryServers(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling query servers")
	return message.ServerList{
		Servers: []block.Piece{
			block.Server{7, "GROUP-SP/", "127.0.0.1", 10887, 0},
			block.Server{6, "SCHE-SP/", "127.0.0.1", 10887, 0},
			block.Server{4, "QUICK0-SP/", "127.0.0.1", 10887, 0},
			block.Server{4, "QUICK1-SP/", "127.0.0.1", 10887, 0},
			block.Server{8, "MENU03-SP/", "127.0.0.1", 12881, 0},
			block.Server{3, "TurboLobas Inc.", "127.0.0.1", 10900, 50},
			block.Server{2, "ACCT03-SP/", "127.0.0.1", 12881, 0},
			block.Server{1, "GATE-SP/", "127.0.0.1", 10887, 0},
		},
	}
}

func HandleRankUrlsQuery(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling rank urls")
	return message.RankUrlList{
		RankUrls: []block.Piece{
			block.RankUrl{0, "http://pes6web.winning-eleven.net/pes6e2/ranking/we10getrank.html"},
			block.RankUrl{1, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10getgrprank.html"},
			block.RankUrl{2, "http://pes6web.winning-eleven.net/pes6e2/ranking/we10RankingWeek.html"},
			block.RankUrl{3, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10GrpRankingWeek.html"},
			block.RankUrl{4, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10RankingCup.html"},
			block.RankUrl{5, "http://www.pes6j.net/server/we10getgrpboard.html"},
			block.RankUrl{6, "http://www.pes6j.net/server/we10getgrpinvitelist.html"},
		},
	}
}

func HandleServerTime(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling server time")
	return message.ServerTime{
		ServerTime: block.ServerTime{Time: time.Now()},
	}
}

func HandleKeepAlive(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling a keep alive")
	return message.KeepAlive{}
}

func HandleDisconnect(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("Handling disconnect")
	return nil
}

func Start() {
	fmt.Println("Here i am the discovery server!")
	s := server.Server{
		FunctionMap: handleMap,
	}
	s.Serve(10881)
}
