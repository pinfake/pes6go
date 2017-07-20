package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

type MenuServer struct {
	logger      *log.Logger
	connections Connections
}

var menuHandlers = map[uint16]Handler{
	0x3080: PlayerFriends,
}

func NewMenuServer() MenuServer {
	return MenuServer{
		logger:      log.New(os.Stdout, "Menu: ", log.LstdFlags),
		connections: NewConnections(),
	}
}

func (s MenuServer) Handlers() map[uint16]Handler {
	return menuHandlers
}

func (s MenuServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s MenuServer) Config() ServerConfig {
	return ServerConfig{
		"serverId": "2",
	}
}

func (s MenuServer) Connections() Connections {
	return s.connections
}

func (s MenuServer) Logger() *log.Logger {
	return s.logger
}

func PlayerFriends(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.NewPlayerFriendsMessage(block.PlayerFriends{})
}

func StartMenu() {
	fmt.Println("Menu Server starting")
	Serve(NewMenuServer(), 12882)
}
