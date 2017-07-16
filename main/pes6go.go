package main

import "fmt"
import (
	"flag"
	"os"

	"github.com/pinfake/pes6go/server/accounting"
	"github.com/pinfake/pes6go/server/discovery"
	"github.com/pinfake/pes6go/server/game"
	"github.com/pinfake/pes6go/server/menu"
)

func main() {
	_ = flag.Bool("d", false, "Run in detached mode")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "\t%s command [arguments]\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "The commands are:\n\n")
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
		go discovery.Start()
		go accounting.Start()
		go menu.Start()
		go game.Start()
		select {}
	case "discovery":
		discovery.Start()
	case "accounting":
		accounting.Start()
	case "menu":
		menu.Start()
	case "game":
		game.Start()
	default:
		flag.Usage()
	}
}
