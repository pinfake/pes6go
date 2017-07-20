package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/storage"
)

type GameServer struct {
	logger      *log.Logger
	connections Connections
}

var gameHandlers = map[uint16]Handler{}

func NewGameServer() GameServer {
	return GameServer{
		logger:      log.New(os.Stdout, "Game: ", log.LstdFlags),
		connections: NewConnections(),
	}
}

func (s GameServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s GameServer) Handlers() map[uint16]Handler {
	return gameHandlers
}

func (s GameServer) Config() ServerConfig {
	return ServerConfig{
		"serverId": "1",
	}
}

func (s GameServer) Connections() Connections {
	return s.connections
}

func (s GameServer) Logger() *log.Logger {
	return s.logger
}

func StartGame() {
	fmt.Println("Game Server starting")
	Serve(NewGameServer(), 10887)
}
