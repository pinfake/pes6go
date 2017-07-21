package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/storage"
)

type GameServer struct {
	config  ServerConfig
	storage storage.Storage
}

var gameHandlers = map[uint16]Handler{}

func NewGameServer() GameServer {
	return GameServer{
		storage: storage.Forged{},
		config: ServerConfig{
			"serverId": "1",
		},
	}
}

func (s GameServer) Storage() storage.Storage {
	return s.storage
}

func (s GameServer) Handlers() map[uint16]Handler {
	return gameHandlers
}

func (s GameServer) Config() ServerConfig {
	return s.config
}

func StartGame() {
	fmt.Println("Game Server starting")
	s := NewServer(log.New(os.Stdout, "Game: ", log.LstdFlags), NewGameServer())
	s.Serve(10887)
}
