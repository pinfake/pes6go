package client

import (
	"net"
	"strconv"
)

type Client struct {
	conn net.Conn
}

func (c Client) Connect(host string, port int) error {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	c.conn = conn
	return err
}

func (c Client) Write(b []byte) {
	c.conn.Write(b)
}
