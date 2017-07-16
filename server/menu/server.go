package menu

import (
	"fmt"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type MenuServer struct {
	storage storage.Storage
}

var handlers = map[uint16]server.Handler{
	0x0003: Disconnect,
	0x0005: KeepAlive,
}

func (s MenuServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func KeepAlive(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling a keep alive")
	return message.KeepAlive{}
}

func Disconnect(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	fmt.Println("Handling disconnect")
	return nil
}

func Start() {
	fmt.Println("Here i am the menu server!")
	server.Serve(MenuServer{
		storage: storage.Forged{},
	}, 12882)
}
