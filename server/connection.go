package server

import (
	"fmt"
	"net"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/network"
)

type Connection struct {
	conn net.Conn
	seq  uint32
	data interface{}
}

func (c Connection) readBlock() (block.Block, error) {
	var data [4096]byte
	slice := data[:]

	n, err := c.conn.Read(slice)
	if err != nil {
		return block.Block{}, err
	}
	got, err := block.ReadBlock(slice[:n])
	if err != nil {
		return block.Block{}, err
	}

	fmt.Printf("READ: % x\n", got)
	return got, nil
}

func (c *Connection) writeMessage(message message.Message) {
	for _, b := range message.GetBlocks() {
		fmt.Printf("Sequence is %d\n", c.seq)
		c.seq++
		b.Header.Sequence = c.seq
		fmt.Printf("WRITE: % x\n", b.GetBytes())
		c.conn.Write(network.Mutate(b.GetBytes()))
	}
}
