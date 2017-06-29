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

type Handler func(block.Block, *Connection) message.Message

type Server interface {
	GetHandlers() map[uint16]Handler
}

func handleConnection(s Server, conn net.Conn) {
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
		m, err := handleBlock(s, b, &c)
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

func handleBlock(s Server, block block.Block, c *Connection) (message.Message, error) {
	method, ok := s.GetHandlers()[block.Header.Query]
	if !ok {
		return nil, fmt.Errorf("Unknown query!")
	}
	return method(block, c), nil
}

func Serve(s Server, port int) {
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
		go handleConnection(s, conn)
	}
}
