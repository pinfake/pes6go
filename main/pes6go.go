package main

import "fmt"
import (
	"flag"
	"os"

	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

var stor storage.Storage

func init() {
	var err error
	stor, err = storage.NewBolt()
	if err != nil {
		panic("Cannot initialize the bolt database: " + err.Error())
	}
}

func main() {
	_ = flag.Bool("d", false, "Run in detached mode")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "\t%s command [arguments]\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "The commands are:\n\n")
		fmt.Fprint(os.Stderr, "\tadmin          Run the administration server")
		fmt.Fprint(os.Stderr, "\tdiscovery      Run a discovery server\n")
		fmt.Fprint(os.Stderr, "\taccounting     Run an accounting server\n")
		fmt.Fprint(os.Stderr, "\tmenu           Run a menu server\n")
		fmt.Fprint(os.Stderr, "\tgame           Run a game server\n")
		fmt.Fprint(os.Stderr, "\tfullhouse      Run all servers at once\n\n")
		fmt.Fprint(os.Stderr, "The arguments are:\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	if len(flag.Args()) < 1 {
		flag.Usage()
		return
	}

	switch args[0] {
	case "fullhouse":
		go server.StartAdmin(stor)
		go server.StartDiscovery()
		go server.StartAccounting(stor)
		go server.StartMenu()
		go server.StartGame()
		select {}
	case "admin":
		server.StartAdmin(stor)
	case "discovery":
		server.StartDiscovery()
	case "accounting":
		server.StartAccounting(stor)
	case "menu":
		server.StartMenu()
	case "game":
		server.StartGame()
	default:
		flag.Usage()
	}
}
