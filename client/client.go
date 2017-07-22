package client

import (
	"net"
	"strconv"
)

type Client struct {
	conn net.Conn
}

func NewClient() Client {
	return Client{}
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

	n, err := c.conn.Read(slice)
	if err != nil {
		return nil, err
	}

	return slice[:n], nil
}

func (c *Client) Close() {
	c.conn.Close()
}
