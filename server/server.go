package server

import (
	"net"
	"fmt"
	"os"
	"strconv"
)

const host = "0.0.0.0"

type Handler interface {
	handleConnection(conn net.Conn)
}

func handleConnection(conn net.Conn, handler Handler) {
	defer conn.Close()
	fmt.Println("Hey!, a connection!")
	handler.handleConnection(conn)
	fmt.Println("It's over!")
}

func Serve(port int, handler Handler) {
	l,err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		go handleConnection(conn, handler)
	}
}

