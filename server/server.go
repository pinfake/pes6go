package server

import (
	"net"
	"fmt"
	"os"
	"time"
)

const host = "0.0.0.0"
const port = "10881"

func handleConnection(conn net.Conn) {
	fmt.Println("Hey!, a connection!")
	for i := 1; i < 6; i++ {
		conn.Write([]byte(fmt.Sprintf("%d\n",i)))
		time.Sleep(1 * time.Second)
	}
	conn.Close()
	fmt.Println("It's over!")
}

func Run() {
	l,err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

//	if l, ok := l.(*net.TCPListener); ok {
//		l.SetDeadline(time.Now().Add(1*time.Second))
//	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

