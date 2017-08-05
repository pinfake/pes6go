package server

import (
	"net"

	"sync"

	"fmt"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/network"
	"github.com/pinfake/pes6go/storage"
)

type Connection struct {
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
	connections map[net.Conn]*Connection
}

func NewConnections() *Connections {
	return &Connections{
		connections: make(map[net.Conn]*Connection),
	}
}

func (conns *Connections) remove(c *Connection) {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	if c.Player != nil {
		fmt.Printf("Going to delete con with player: %s\n", c.Player.Name)
	}
	delete(conns.connections, c.conn)
}

func (conns *Connections) add(c net.Conn) *Connection {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	connection := Connection{
		seq:  0,
		conn: c,
	}
	conns.connections[c] = &connection
	return &connection
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
	fmt.Printf("Sending to lobby: %d\n", lobbyId)
	for _, conn := range conns.connections {
		fmt.Printf("Found player: %s in lobby: %d\n",
			conn.Player.Name, conn.LobbyId)
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
