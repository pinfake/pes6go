package client

import (
	"net"
	"strconv"

	"time"

	"github.com/pinfake/pes6go/crypt"
	"github.com/pinfake/pes6go/data/block"
)

type Client struct {
	seq  uint32
	conn net.Conn
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(host string, port int) error {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	c.conn = conn
	return err
}

func (c *Client) Write(b []byte) {
	c.conn.Write(b)
}

func (c *Client) Read() ([]byte, error) {
	var data [4096]byte
	slice := data[:]

	c.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	n, err := c.conn.Read(slice)
	if err != nil {
		return nil, err
	}

	return slice[:n], nil
}

func (c *Client) WriteBlock(b *block.Block) {

	c.seq++
	b.Sign(c.seq)
	c.conn.Write(crypt.ApplyMask(b.GetBytes()))
}

func (c *Client) Close() {
	c.conn.Close()
}
