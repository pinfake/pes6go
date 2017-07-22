package server

import (
	"log"
	"testing"

	"os"

	"github.com/pinfake/pes6go/client"
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

func connect(c *client.Client, t *testing.T) {
	err := c.Connect("localhost", port)
	if err != nil {
		t.Error("Error connecting: %s", err.Error())
	}
}

func TestConnect(t *testing.T) {
	t.Run("Should be able to connect", func(t *testing.T) {
		c := client.NewClient()
		connect(&c, t)
		c.Close()
	})
}

func TestSendInvalidData(t *testing.T) {
	t.Run("Shouldn't crash on invalid data", func(t *testing.T) {
		c := client.Client{}
		connect(&c, t)
		c.Write([]byte{0x01, 0x02, 0x03})
		c.Read()
		c.Close()
		connect(&c, t)
		c.Close()
	})
}

func TestSendProperHeadShorterBody(t *testing.T) {

}

func TestSendProperHeadLongerBody(t *testing.T) {

}

func TestSendMoreThanReadBuffer(t *testing.T) {

}

func TestSendUnknownQuery(t *testing.T) {

}
