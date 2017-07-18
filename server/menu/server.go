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
	0x3080: PlayerFriends,
}

func PlayerFriends(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewPlayerFriendsMessage(block.PlayerFriends{})
}

func (s MenuServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func (s MenuServer) GetStorage() storage.Storage {
	return storage.Forged{}
}

func Start() {
	fmt.Println("Menu Server starting")
	server.Serve(MenuServer{}, 12882)
}
