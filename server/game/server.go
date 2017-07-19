package game

import (
	"fmt"

	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type GameServer struct {
	data GameServerData
}

type GameServerData struct {
	Hola  string
	Adios int
}

var handlers = map[uint16]server.Handler{}

func (s GameServer) GetStorage() storage.Storage {
	return storage.Forged{}
}

func (s GameServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func (s GameServer) GetConfig() server.ServerConfig {
	return server.ServerConfig{
		"serverId": "1",
	}
}

func (s GameServer) GetData() interface{} {
	return s.data
}

func Start() {
	fmt.Println("Game Server starting")
	server.Serve(GameServer{}, 10887)
}
