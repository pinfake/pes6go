package server

import (
	"log"
	"testing"

	"os"

	"crypto/rand"

	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/storage"
)

const port = 19780

var s *Server

type emptyServer struct {
}

func (s emptyServer) Config() ServerConfig {
	return nil
}

func (s emptyServer) Storage() storage.Storage {
	return nil
}

func (s emptyServer) Handlers() map[uint16]Handler {
	return map[uint16]Handler{}
}

func NewEmptyServer() *Server {
	return NewServer(
		log.New(os.Stdout, "test: ", log.LstdFlags),
		emptyServer{},
	)
}

func init() {
	s = NewEmptyServer()
	go s.Serve(port)
}

func getRandom(size int) []byte {
	data := make([]byte, size)
	rand.Read(data)

	return data
}

func craftBlock(query uint16, size uint16, data []byte) *block.Block {
	b := block.Block{
		Header: block.Header{
			Query:    query,
			Size:     size,
			Sequence: 0,
			Hash:     [16]byte{},
		},
		Body: block.GenericBody{
			Data: data,
		},
	}

	return &b
}

func connect(c *client.Client, t *testing.T) {
	err := c.Connect("localhost", port)
	if err != nil {
		t.Error("Error connecting: %s", err.Error())
	}
}

func TestConnect(t *testing.T) {
	t.Run("Should be able to connect", func(t *testing.T) {
		c := client.NewClient()
		connect(c, t)
		c.Close()
	})
}

func TestSendInvalidData(t *testing.T) {
	t.Run("Shouldn't crash on invalid data", func(t *testing.T) {
		c := client.NewClient()
		connect(c, t)
		c.Write([]byte{0x01, 0x02, 0x03})
		c.Read()
		c.Close()
	})
}

func TestSendProperHeadLongerBody(t *testing.T) {
	t.Run("Shouldn't crash on longer body that header says", func(t *testing.T) {
		b := craftBlock(0x3001, 10, getRandom(100))
		c := client.NewClient()
		connect(c, t)
		c.WriteBlock(b)
		c.Read()
		c.Close()
	})
}

func TestSendProperHeadShorterBody(t *testing.T) {
	t.Run("Shouldn't crash on shorter body that header says", func(t *testing.T) {
		b := craftBlock(0x3001, 100, getRandom(10))
		c := client.NewClient()
		connect(c, t)
		c.WriteBlock(b)
		c.Read()
		c.Close()
	})
}

func TestSendMoreThanReadBuffer(t *testing.T) {
	t.Run("Shouldn't crash on more data than actual buffer", func(t *testing.T) {
		c := client.NewClient()
		connect(c, t)
		c.Write(getRandom(10000))
		c.Read()
		c.Close()
	})
}

func TestSendUnknownQuery(t *testing.T) {
	t.Run("Shouldn't crash on unknown query", func(t *testing.T) {
		b := craftBlock(0x1234, 100, getRandom(100))
		c := client.NewClient()
		connect(c, t)
		c.WriteBlock(b)
		c.Read()
		c.Close()
	})
}
