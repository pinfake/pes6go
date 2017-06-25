package accounting

import (
	"fmt"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
)

var handleMap = map[uint16]func(block.Block, *server.Connection) message.Message{
	0x3001: HandleInit,
	0x3003: HandleLogin,
	0x3010: HandleProfiles,
	0x0005: HandleKeepAlive,
	0x0003: HandleDisconnect,
}

func HandleProfiles(_ block.Block, _ *server.Connection) message.Message {
	return nil
}

func HandleLogin(b block.Block, _ *server.Connection) message.Message {
	auth := block.NewAthentication(b)
	fmt.Println("I am handling login")
	fmt.Printf("key: % x\n", auth.Key)
	fmt.Printf("password: % x\n", auth.Password)
	fmt.Printf("unknown: % x\n", auth.Unknown)
	fmt.Printf("roster: % x\n", auth.RosterHash)
	return message.LoginResponse{}
}

func HandleInit(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling init")
	return message.AccountingInit{}
}

func HandleKeepAlive(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling a keep alive")
	return message.KeepAlive{}
}

func HandleDisconnect(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("Handling disconnect")
	return nil
}

func Start() {
	fmt.Println("Here i am the accounting server!")
	s := server.Server{
		FunctionMap: handleMap,
	}
	s.Serve(12881)
}
