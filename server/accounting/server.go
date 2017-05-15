package accounting

import (
	"fmt"

	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
	"github.com/pinfake/pes6go/network/messages/accounting"
	"github.com/pinfake/pes6go/network/messages/common"
	"github.com/pinfake/pes6go/server"
)

var handleMap = map[uint16]func(blocks.Block, *server.Connection) messages.Message{
	0x3001: HandleInit,
	0x3003: HandleLogin,
	0x0005: HandleKeepAlive,
	0x0003: HandleDisconnect,
}

type Server struct {
	server.Handler
}

func HandleLogin(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling login")
	return nil
}

func HandleInit(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling init")
	return accounting.InitMessage{}
}

func HandleKeepAlive(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("I am handling a keep alive")
	return common.KeepAlive{}
}

func HandleDisconnect(_ blocks.Block, _ *server.Connection) messages.Message {
	fmt.Println("Handling disconnect")
	return nil
}

func (s Server) HandleBlock(block blocks.Block, c *server.Connection) (messages.Message, error) {
	method, ok := handleMap[block.Header.Query]
	if !ok {
		return nil, fmt.Errorf("Unknown query!")
	}
	return method(block, c), nil
}

func Start() {
	fmt.Println("Here i am the accounting server!")
	s := Server{}
	server.Serve(12881, s)
}
