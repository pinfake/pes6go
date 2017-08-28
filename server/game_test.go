package server

import (
	"testing"

	"log"
	"os"

	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/storage"
)

const gsTestPort = 10887

func init() {
	stor, err := storage.NewBolt()
	if err != nil {
		panic("Cannot initialize the bolt database: " + err.Error())
	}
	s := NewServer(log.New(os.Stdout, "Game: ", log.LstdFlags), NewGameServerHandler(stor))
	go s.Serve(gsTestPort)
	go StartGame(stor)
}

func TestGameConnect(t *testing.T) {
	t.Run("Should connect", func(t *testing.T) {
		c := client.NewClient()
		connect(c, gsTestPort, t)
		c.Close()
	})
}

func TestGameLogin(t *testing.T) {
	t.Run("Should login", func(t *testing.T) {
		au := client.NewAutomatedClient("localhost", gsTestPort,
			[]client.CommandHandler{
				client.Init{},
				client.Login{
					Command: client.Command{
						Data: map[string]interface{}{
							"key":      "NSTXC54CKP3KLLL26M65",
							"password": "pin-qwe",
						},
					},
				},
			},
		)
		au.Run()
	})
}
