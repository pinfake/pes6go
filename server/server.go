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
	HandleBlock(block blocks.Block) messages.Message
}

type Connection struct {
	conn net.Conn
}

func (c Connection) readBlock() (blocks.Block, error) {
	var data [4096]byte
	slice := data[:]

	n, err := c.conn.Read(slice)
	if err != nil {
		return blocks.Block{}, err
	}
	got, err := blocks.ReadBlock(slice[:n])
	if err != nil {
		return blocks.Block{}, err
	}
	//got := network.Mutate(slice[:n])
	fmt.Printf("% x\n", got)
	return got, nil
}

func (c Connection) writeMessage(message messages.Message) {
	fmt.Println("I should write something here")
}

func handleConnection(conn net.Conn, handler Handler) {
	defer conn.Close()
	c := Connection{conn: conn}
	fmt.Println("Hey!, a connection!")
	block, err := c.readBlock()
	if err != nil {
		panic("Couldn't properly read")
	}
	message := handler.HandleBlock(block)
	bs := message.GetBlocks()
	fmt.Printf("Going to write: % x", bs)
	c.writeMessage(message)
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
