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
	0x0005: HandleKeepAlive,
	0x0003: HandleDisconnect,
}

func HandleLogin(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling login")
	return nil
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
