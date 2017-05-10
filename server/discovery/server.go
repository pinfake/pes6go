package discovery

import (
	"fmt"
	"github.com/pinfake/pes6go/server"
	"net"
	"time"
)

type Server struct {
	server.Handler
}

func (s Server) HandleConnection (conn net.Conn) {
	for i := 1; i < 6; i++ {
		conn.Write([]byte(fmt.Sprintf("%d\n",i)))
		time.Sleep(1 * time.Second)
	}
}

func Start() {
	fmt.Println("Here i am the discovery server!")
	discovery := Server{}
	server.Serve(10881, discovery)
}
