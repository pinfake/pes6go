package server

import (
	"fmt"
	"net"
	"strconv"

	"log"

	"time"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

const host = "0.0.0.0"

type Handler func(*Server, *block.Block, *Connection) message.Message

type ServerConfig map[string]string

type Server struct {
	logger      *log.Logger
	connections *Connections
	listener    net.Listener
	ServerHandler
}

type ServerHandler interface {
	Handlers() map[uint16]Handler
	Storage() storage.Storage
	Config() ServerConfig
}

func (s *Server) Log(c *Connection, format string, v ...interface{}) {
	prefix := ""
	if c != nil {
		prefix += c.conn.RemoteAddr().String() + " " + strconv.Itoa(c.id) + " "
	}
	s.logger.Printf(prefix+format, v...)
}

func (s *Server) handleConnection(conn net.Conn) error {
	c := s.connections.add(conn)
	defer s.connections.remove(c.id)
	s.Log(c, "Incoming connection")
	for {
		b, err := c.readBlock()
		if err != nil {
			s.Log(c, "Cannot read block: %s", err)
			return fmt.Errorf("Cannot read block: %s", err)
		}
		s.Log(c, "R <- %x", b)
		m, err := s.handleBlock(b, c)
		if err != nil {
			s.Log(c, "handleConnection: %s", err)
			return fmt.Errorf("handleConnection: %s", err)
		}
		if m == nil {
			break
		}
		bs := m.GetBlocks()
		s.Log(c, "W -> %x", bs)
		c.writeMessage(m)
	}
	s.Log(c, "Closing connection")
	return nil
}

func (s *Server) handleBlock(block *block.Block, c *Connection) (message.Message, error) {
	var method, ok = s.Handlers()[block.Header.Query]
	if !ok {
		method, ok = handlers[block.Header.Query]
		if !ok {
			return nil, fmt.Errorf("Unknown query!")
		}
	}
	return method(s, block, c), nil
}

func (s *Server) Serve(port int) {
	l, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		s.Log(nil, "Error listening: %s", err.Error())
		panic("Error listening:" + err.Error())
	}
	s.listener = l
	s.Log(nil, "Server listening on "+host+":"+strconv.Itoa(port))

	defer s.Shutdown()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.Log(nil, "Error accepting: ", err)
			return
		}
		go s.handleConnection(conn)
	}
}

func NewServer(logger *log.Logger, handler ServerHandler) *Server {
	s := Server{
		logger:        logger,
		connections:   NewConnections(),
		ServerHandler: handler,
	}
	return &s
}

func (s Server) Shutdown() {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s Server) WaitUntilDone() {
	time.Sleep(10000 * time.Millisecond)
}
