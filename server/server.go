package server

import (
	"fmt"
	"net"
	"strconv"

	"log"

	"encoding/json"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/data/types"
	"github.com/pinfake/pes6go/storage"
)

const host = "0.0.0.0"

type Handler func(*Server, *block.Block, *Connection) message.Message

type ServerConfig map[string]string

type Server struct {
	logger      *log.Logger
	connections *types.IdMap
	listener    net.Listener
	lobbies     []*block.Lobby
	ServerHandler
}

type ServerHandler interface {
	Handlers() map[uint16]Handler
	Storage() storage.Storage
	Config() ServerConfig
	Data() interface{}
}

func (s *Server) Log(c *Connection, format string, v ...interface{}) {
	prefix := ""
	if c != nil {
		prefix += c.conn.RemoteAddr().String() + " "
	}
	s.logger.Printf(prefix+format, v...)
}

func (s *Server) handleConnection(conn net.Conn) error {
	c := &Connection{
		id:      s.connections.GetNewId(),
		LobbyId: 0xff, // 0xff meaning no lobby
		seq:     0,
		conn:    conn,
		logger:  s.logger,
	}
	s.connections.Add(c.id, c)
	defer s.closeConnection(c)
	s.Log(c, "Incoming connection")
	for {
		b, err := c.readBlock()
		if err != nil {
			Disconnect(s, nil, c)
			s.Log(c, "Cannot read block: %s", err)
			return fmt.Errorf("Cannot read block: %s", err)
		}
		m, err := s.handleBlock(b, c)
		if err != nil {
			s.Log(c, "handleConnection: %s", err)
			return fmt.Errorf("handleConnection: %s", err)
		}
		if m == nil {
			continue
		}
		c.writeMessage(m)
	}
	return nil
}

func (s *Server) closeConnection(c *Connection) {
	s.Log(c, "Closing connection")
	s.connections.Delete(c.id)
	if c.conn != nil {
		c.conn.Close()
	}
}

func (s *Server) handleBlock(block *block.Block, c *Connection) (message.Message, error) {
	var method, ok = s.Handlers()[block.Header.Query]
	if !ok {
		method, ok = handlers[block.Header.Query]
		if !ok {
			Disconnect(s, block, c)
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
			s.Log(nil, "Error accepting: %s", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) initializeLobbies() {
	if lobbies, ok := s.Config()["lobbies"]; ok {
		err := json.Unmarshal([]byte(lobbies), &s.lobbies)
		if err != nil {
			panic(err)
		}
	}
	for _, lobby := range s.lobbies {
		lobby.Rooms = types.NewIdMap()
	}
}

func NewServer(logger *log.Logger, handler ServerHandler) *Server {
	s := Server{
		logger:        logger,
		connections:   types.NewIdMap(),
		ServerHandler: handler,
	}
	s.initializeLobbies()
	return &s
}

func (s *Server) Shutdown() {
	if s.listener != nil {
		s.listener.Close()
	}
}
