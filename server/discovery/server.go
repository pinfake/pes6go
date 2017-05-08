package discovery

import (
	"fmt"
	"github.com/pinfake/pes6go/server"
)

func Start() {
	fmt.Println("Here i am the discovery server!")
	server.Run()
}
