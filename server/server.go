package server

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"log"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

const host = "0.0.0.0"

type Handler func(Server, block.Block, *Connection) message.Message

type ServerConfig map[string]string

type Server interface {
	Logger() *log.Logger
	Connections() Connections
	Config() ServerConfig
	Storage() storage.Storage
	Handlers() map[uint16]Handler
}

func Log(s Server, c *Connection, format string, v ...interface{}) {
	prefix := c.conn.RemoteAddr().String() + " " + strconv.Itoa(c.id) + " "
	s.Logger().Printf(prefix+format, v...)
}

func handleConnection(s Server, conn net.Conn) {
	c := s.Connections().add(conn)
	defer s.Connections().remove(c.id)
	Log(s, c, "Incoming connection")
	for {
		b, err := c.readBlock()
		if err != nil {
			panic("Couldn't properly read " + err.Error())
		}
		Log(s, c, "R <- %x", b)
		m, err := handleBlock(s, b, c)
		if err != nil {
			panic(err)
		}
		if m == nil {
			break
		}
		bs := m.GetBlocks()
		Log(s, c, "W -> %x", bs)
		c.writeMessage(m)
	}
	Log(s, c, "Closing connection")
}

func handleBlock(s Server, block block.Block, c *Connection) (message.Message, error) {
	var method, ok = s.Handlers()[block.Header.Query]
	if !ok {
		method, ok = handlers[block.Header.Query]
		if !ok {
			return nil, fmt.Errorf("Unknown query!")
		}
	}
	return method(s, block, c), nil
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
