package game

import (
	"fmt"

	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type GameServer struct {
	connections server.Connections
}

var handlers = map[uint16]server.Handler{}

func NewGameServer() GameServer {
	return GameServer{connections: server.NewConnections()}
}

func (s GameServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s GameServer) Handlers() map[uint16]server.Handler {
	return handlers
}

func (s GameServer) Config() server.ServerConfig {
	return server.ServerConfig{
		"serverId": "1",
	}
}

func (s GameServer) Connections() server.Connections {
	return s.connections
}

func Start() {
	fmt.Println("Game Server starting")
	server.Serve(NewGameServer(), 10887)
}
