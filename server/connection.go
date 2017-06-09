package server

import (
	"net"
	"github.com/pinfake/pes6go/network/blocks"
	"fmt"
	"github.com/pinfake/pes6go/network/messages"
	"github.com/pinfake/pes6go/network"
)

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