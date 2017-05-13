package discovery

import (
	"fmt"
	"net"
	"time"

	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
	"github.com/pinfake/pes6go/network/messages/discovery"
	"github.com/pinfake/pes6go/server"
)

var handleMap map[uint16]func(block blocks.Block) messages.Message = map[uint16]func(block blocks.Block) messages.Message{
	0x2005: HandleDiscoveryInit,
}

type Server struct {
	server.Handler
}

type VoidMessage struct {
	messages.Message
}

func (m VoidMessage) getBlocks() []blocks.Block {
	return []blocks.Block{}
}

func HandleDiscoveryInit(block blocks.Block) messages.Message {
	fmt.Println("I am handling discovery init")
	return discovery.Response{}
}

func (s Server) HandleBlock(block blocks.Block) messages.Message {
	return handleMap[block.Header.Query](block)
}

func (s Server) HandleConnection(conn net.Conn) {
	for i := 1; i < 6; i++ {
		conn.Write([]byte(fmt.Sprintf("%d\n", i)))
		time.Sleep(1 * time.Second)
	}
}

func Start() {
	fmt.Println("Here i am the s server!")
	s := Server{}
	server.Serve(10881, s)
}
