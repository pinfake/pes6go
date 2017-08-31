package server

import (
	"net"

	"fmt"

	"github.com/pinfake/pes6go/crypt"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/data/types"
	"github.com/pinfake/pes6go/storage"
	"log"
)

type Connection struct {
	id      uint32
	conn    net.Conn
	seq     uint32
	logger  *log.Logger
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
	c.log("R <- %v", got)

	return got, nil
}

func (c *Connection) writeMessage(message message.Message) {
	for _, b := range message.GetBlocks() {
		c.writeBlock(b)
	}
}

func (c *Connection) writeBlock(b *block.Block) {
	c.seq++
	b.Sign(c.seq)
	c.log("W -> %v", b)
	c.conn.Write(crypt.ApplyMask(b.GetBytes()))
}

func findByPlayerId(idmap *types.IdMap, playerId uint32) *Connection {
	defer idmap.RUnlock()
	idmap.RLock()
	for _, e := range idmap.Data {
		c := e.(*Connection)
		if c.Player == nil {
			continue
		}
		if c.Player.Id == playerId {
			return c
		}
	}
	return nil
}

func playersInLobby(idmap *types.IdMap, lobbyId byte) []*block.Player {
	var ret []*block.Player
	defer idmap.RUnlock()
	idmap.RLock()
	for _, e := range idmap.Data {
		c := e.(*Connection)
		if c.LobbyId == lobbyId {
			ret = append(ret, c.Player)
		}
	}
	return ret
}

func sendToLobby(idmap *types.IdMap, lobbyId byte, m message.Message) {
	defer idmap.RUnlock()
	idmap.RLock()
	fmt.Printf("Sending to lobby %v\n", m)
	for _, e := range idmap.Data {
		c := e.(*Connection)
		if c.LobbyId == lobbyId {
			c.writeMessage(m)
		}
	}
}

func sendToRoom(idmap *types.IdMap, roomId uint32, m message.Message, me *Connection) {
	defer idmap.RUnlock()
	idmap.RLock()
	for _, e := range idmap.Data {
		c := e.(*Connection)
		if c == me {
			continue
		}
		if c.Player != nil && c.Player.RoomId == roomId {
			c.writeMessage(m)
		}
	}
}

func (c *Connection) log(format string, v ...interface{}) {
	prefix := ""
	if c != nil {
		prefix += c.conn.RemoteAddr().String() + " "
	}
	c.logger.Printf(prefix+format, v...)
}
