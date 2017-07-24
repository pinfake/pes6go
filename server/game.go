package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

type GameServerData struct {
	rooms []*block.Room
}

type GameServer struct {
	data    GameServerData
	config  ServerConfig
	storage storage.Storage
}

var gameHandlers = map[uint16]Handler{
	//	0x4102:
	0x4210: PlayersInLobby,
	0x4300: RoomsInLobby,
}

func NewGameServerHandler() GameServer {
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

func (s GameServer) Data() interface{} {
	return s.data
}

func PlayersInLobby(s *Server, _ *block.Block, c *Connection) message.Message {
	return message.NewPlayersInLobbyMessage(
		s.connections.playersInLobby(c.LobbyId),
	)
}

func RoomsInLobby(s *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewRoomsInLobbyMessage(
		s.Data().(GameServerData).rooms,
	)
}

func StartGame() {
	fmt.Println("Game Server starting")
	s := NewServer(log.New(os.Stdout, "Game: ", log.LstdFlags), NewGameServerHandler())
	s.Serve(10887)
}
