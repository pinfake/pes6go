package server

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
)

const host = "0.0.0.0"

type Server struct {
	Handler
	FunctionMap map[uint16]func(block.Block, *Connection) message.Message
}

type Handler interface {
	HandleBlock(block block.Block, c *Connection) (message.Message, error)
}

func (s Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	c := Connection{
		conn: conn,
		seq:  0,
	}
	fmt.Println("Hey!, a connection!")
	for {
		b, err := c.readBlock()
		if err != nil {
			panic("Couldn't properly read")
		}
		m, err := s.HandleBlock(b, &c)
		if err != nil {
			panic(err)
		}
		if m == nil {
			break
		}
		bs := m.GetBlocks()
		fmt.Printf("Going to write: % x", bs)
		c.writeMessage(m)
	}
	fmt.Println("It's over!")
}

func (s Server) HandleBlock(block block.Block, c *Connection) (message.Message, error) {
	method, ok := s.FunctionMap[block.Header.Query]
	if !ok {
		return nil, fmt.Errorf("Unknown query!")
	}
	return method(block, c), nil
}

func (s Server) Initialize(f map[uint16]func(block.Block, *Connection) message.Message) {
	s.FunctionMap = f
}

func (s Server) Serve(port int) {
	l, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		go s.handleConnection(conn)
	}
}
