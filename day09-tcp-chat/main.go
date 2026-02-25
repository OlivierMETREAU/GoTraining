package main

import (
	"fmt"
	"net"

	"example.com/day09-tcp-chat/tcp"
)

func main() {
	fmt.Println("day09-tcp-chat")

	srv := tcp.NewServer()

	// Start the server event loop
	stop := make(chan struct{})
	go srv.Run(stop)

	// Start the TCP listener
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on :9000")

	for {
		_, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		// This will be implemented in client.go
		//go tcp.HandleClient(srv, conn)
	}
}
