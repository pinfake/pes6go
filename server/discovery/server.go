package discovery

import (
	"fmt"
	"time"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type DiscoveryServer struct {
	storage storage.Storage
}

var handlers = map[uint16]server.Handler{
	0x2008: Init,
	0x2006: ServerTime,
	0x2005: Servers,
	0x2200: RankUrls,
	0x0005: KeepAlive,
	0x0003: Disconnect,
}

func (s DiscoveryServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func Init(s server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewServerNewsMessage(
		s.(DiscoveryServer).storage.GetServerNews(),
	)
}

func Servers(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewServerListMessage(
		[]block.Server{
			{7, "GROUP-SP/", "127.0.0.1", 10887, 0},
			{6, "SCHE-SP/", "127.0.0.1", 10887, 0},
			{4, "QUICK0-SP/", "127.0.0.1", 10887, 0},
			{4, "QUICK1-SP/", "127.0.0.1", 10887, 0},
			{8, "MENU03-SP/", "127.0.0.1", 12882, 0},
			{3, "TurboLobas Inc.", "127.0.0.1", 10900, 50},
			{2, "ACCT03-SP/", "127.0.0.1", 12881, 0},
			{1, "GATE-SP/", "127.0.0.1", 10887, 0},
		},
	)
}

func RankUrls(s server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewRankUrlListMessage(
		s.(DiscoveryServer).storage.GetRankUrls(),
	)
}

func ServerTime(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.ServerTime{
		ServerTime: block.ServerTime{Time: time.Now()},
	}
}

func KeepAlive(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.KeepAlive{}
}

func Disconnect(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return nil
}

func Start() {
	fmt.Println("Discovery Server starting")
	server.Serve(DiscoveryServer{
		storage: storage.Forged{},
	}, 10881)
}
