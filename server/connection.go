package server

import (
	"net"

	"sync"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/network"
	"github.com/pinfake/pes6go/storage"
)

type Connection struct {
	id      int
	conn    net.Conn
	seq     uint32
	Account *storage.Account
	LobbyId byte
	RoomId  uint32
	Player  *block.Player
}

func (c *Connection) readBlock() (*block.Block, error) {
	var data [4096]byte
	slice := data[:]

	n, err := c.conn.Read(slice)
	if err != nil {
		return nil, err
	}
	got, err := block.ReadBlock(slice[:n])
	if err != nil {
		return nil, err
	}

	return got, nil
}

func (c *Connection) writeMessage(message message.Message) {
	for _, b := range message.GetBlocks() {
		c.seq++
		b.Header.Sequence = c.seq
		c.conn.Write(network.Mutate(b.GetBytes()))
	}
}

type Connections struct {
	mu          sync.RWMutex
	connections map[int]*Connection
}

func NewConnections() *Connections {
	return &Connections{
		connections: make(map[int]*Connection),
	}
}

func (conns *Connections) remove(id int) {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	delete(conns.connections, id)
}

func (conns *Connections) add(c net.Conn) *Connection {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	connection := Connection{
		id:   conns.newId(),
		seq:  0,
		conn: c,
	}
	conns.connections[connection.id] = &connection
	return &connection
}

func (conns *Connections) newId() int {
	return len(conns.connections) + 1
}

func (conns *Connections) findByPlayerId(id uint32) *Connection {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	return nil
}

func (conns *Connections) findByPlayerName(name string) *Connection {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	return nil
}

func (conns *Connections) playersInLobby(lobbyId byte) []*block.Player {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	var ret []*block.Player
	for _, conn := range conns.connections {
		if conn.LobbyId == lobbyId {
			ret = append(ret, conn.Player)
		}
	}
	return ret
}

func (conns *Connections) sendToLobby(lobbyId byte, m message.Message) {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	for _, conn := range conns.connections {
		if conn.LobbyId == lobbyId {
			conn.writeMessage(m)
		}
	}
}

func (conns *Connections) sendToRoom(roomId uint32, m message.Message) {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	for _, conn := range conns.connections {
		if conn.RoomId == roomId {
			conn.writeMessage(m)
		}
	}
}
