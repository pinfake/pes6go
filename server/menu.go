package server

import (
	"fmt"

	"log"

	"os"

	"github.com/pinfake/pes6go/storage"
)

type MenuServer struct {
	config  ServerConfig
	storage storage.Storage
}

var menuHandlers = map[uint16]Handler{}

func (s MenuServer) Handlers() map[uint16]Handler {
	return menuHandlers
}

func (s MenuServer) Storage() storage.Storage {
	return s.storage
}

func (s MenuServer) Config() ServerConfig {
	return s.config
}

func (s MenuServer) Data() interface{} {
	return nil
}

func NewMenuServerHandler(stor storage.Storage) MenuServer {
	return MenuServer{
		storage: stor,
		config: ServerConfig{
			"serverId": "2",
		},
	}
}

func StartMenu(stor storage.Storage) {
	fmt.Println("Menu Server starting")
	s := NewServer(log.New(os.Stdout, "Menu: ", log.LstdFlags), NewMenuServerHandler(stor))
	s.Serve(12882)
}
