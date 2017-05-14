package server

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/pinfake/pes6go/network"
	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
)

const host = "0.0.0.0"

type Handler interface {
	HandleBlock(block blocks.Block) (messages.Message, error)
}

type Connection struct {
	conn net.Conn
	seq  uint16
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
	//c.seq = got.Header.Sequence
	//got := network.Mutate(slice[:n])
	fmt.Printf("% x\n", got)
	return got, nil
}

func (c *Connection) writeMessage(message messages.Message) {
	fmt.Println("I should write something here")
	for _, block := range message.GetBlocks() {
		fmt.Printf("Seq vale %d\n", c.seq)
		c.seq++
		block.Header.Sequence = c.seq
		fmt.Printf("% x\n", block.GetBytes())
		c.conn.Write(network.Mutate(block.GetBytes()))
	}
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
		message, err := handler.HandleBlock(block)
		if err != nil {
			panic(err)
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
