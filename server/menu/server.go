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
	0x3001: Init,
	0x3003: Login,
	0x4100: Capabilities,
}

func Capabilities(_ server.Server, b block.Block, _ *server.Connection) message.Message {
	_ = block.NewPlayerSelected(b)

	return nil
}

func Init(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.AccountingInit{}
}

func Login(_ server.Server, b block.Block, _ *server.Connection) message.Message {
	auth := block.NewAthentication(b)
	fmt.Println("I am handling login")
	fmt.Printf("key: % x\n", auth.Key)
	fmt.Printf("password: % x\n", auth.Password)
	fmt.Printf("unknown: % x\n", auth.Unknown)
	fmt.Printf("roster: % x\n", auth.RosterHash)
	return message.LoginResponse{
		block.Ok,
	}
}

func (s MenuServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func KeepAlive(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.KeepAlive{}
}

func Disconnect(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return nil
}

func Start() {
	fmt.Println("Menu Server starting")
	server.Serve(MenuServer{
		storage: storage.Forged{},
	}, 12882)
}
