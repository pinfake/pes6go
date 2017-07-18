package game

import (
	"fmt"

	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type GameServer struct {
}

var handlers = map[uint16]server.Handler{}

func (s GameServer) GetStorage() storage.Storage {
	return storage.Forged{}
}

func (s GameServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func Start() {
	fmt.Println("Game Server starting")
	server.Serve(GameServer{}, 10887)
}
