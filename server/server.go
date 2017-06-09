package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
)

const host = "0.0.0.0"

type Handler interface {
	HandleBlock(block blocks.Block, c *Connection) (messages.Message, error)
}

func handleConnection(conn net.Conn, handler Handler) {
	defer conn.Close()
	c := Connection{
		conn: conn,
		seq:  0,
	}
	fmt.Println("Hey!, a connection!")
	for {
		block, err := c.readBlock()
		if err != nil {
			panic("Couldn't properly read")
		}
		message, err := handler.HandleBlock(block, &c)
		if err != nil {
			panic(err)
		}
		if message == nil {
			break
		}
		bs := message.GetBlocks()
		fmt.Printf("Going to write: % x", bs)
		c.writeMessage(message)
	}
	//handler.HandleConnection(conn)
	fmt.Println("It's over!")
}

func Serve(port int, handler Handler) {
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
		go handleConnection(conn, handler)
	}
}