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
	config  ServerConfig
	storage storage.Storage
}

var menuHandlers = map[uint16]Handler{
	0x3080: PlayerFriends,
}

func (s MenuServer) Handlers() map[uint16]Handler {
	return menuHandlers
}

func (s MenuServer) Storage() storage.Storage {
	return s.storage
}

func (s MenuServer) Config() ServerConfig {
	return s.config
}

func NewMenuServer() MenuServer {
	return MenuServer{
		storage: storage.Forged{},
		config: ServerConfig{
			"serverId": "2",
		},
	}
}

func PlayerFriends(_ *Server, _ block.Block, _ *Connection) message.Message {
	return message.NewPlayerFriendsMessage(block.PlayerFriends{})
}

func StartMenu() {
	fmt.Println("Menu Server starting")
	s := NewServer(log.New(os.Stdout, "Menu: ", log.LstdFlags), NewMenuServer())
	s.Serve(12882)
}
