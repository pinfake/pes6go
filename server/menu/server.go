package menu

import (
	"fmt"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type MenuServer struct {
	connections server.Connections
}

var handlers = map[uint16]server.Handler{
	0x3080: PlayerFriends,
}

func NewMenuServer() MenuServer {
	return MenuServer{connections: server.NewConnections()}
}

func (s MenuServer) Handlers() map[uint16]server.Handler {
	return handlers
}

func (s MenuServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s MenuServer) Config() server.ServerConfig {
	return server.ServerConfig{
		"serverId": "2",
	}
}

func (s MenuServer) Connections() server.Connections {
	return s.connections
}

func PlayerFriends(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewPlayerFriendsMessage(block.PlayerFriends{})
}

func Start() {
	fmt.Println("Menu Server starting")
	server.Serve(NewMenuServer(), 12882)
}
