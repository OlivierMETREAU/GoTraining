package main

import (
	"fmt"
	"net"
	"time"

	"example.com/day09-tcp-chat/tcp"
)

type simpleAddr string

func (a simpleAddr) Network() string { return "mock" }
func (a simpleAddr) String() string  { return string(a) }

type namedConn struct {
	net.Conn
	addr net.Addr
}

func (n namedConn) RemoteAddr() net.Addr {
	return n.addr
}

// createMockClient attaches a net.Conn to the server using HandleClient
func createMockClient(s *tcp.Server, name string) (net.Conn, net.Conn) {
	serverSide, clientSide := net.Pipe()

	// Wrap serverSide with a unique RemoteAddr
	namedServerConn := namedConn{
		Conn: serverSide,
		addr: simpleAddr(name), // unique per client
	}

	go func() {
		tcp.HandleClient(s, namedServerConn)
	}()

	// Set nickname
	fmt.Fprintf(clientSide, "/nick %s\n", name)

	return namedServerConn, clientSide
}

func main() {
	fmt.Println("day09-tcp-chat")

	srv := tcp.NewServer()

	// Start server event loop
	stop := make(chan struct{})
	go srv.Run(stop)

	// Create two mock clients
	_, c1 := createMockClient(srv, "Alice")
	_, c2 := createMockClient(srv, "Bob")

	// Reader goroutines to print what each client receives
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := c1.Read(buf)
			if err != nil {
				return
			}
			fmt.Print("Alice receives: ", string(buf[:n]))
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := c2.Read(buf)
			if err != nil {
				return
			}
			fmt.Print("Bob receives: ", string(buf[:n]))
		}
	}()

	// Broadcast 10 messages
	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("Message %d from server", i)
		srv.Broadcast <- msg
		time.Sleep(200 * time.Millisecond)
	}

	// Give time for messages to flush
	time.Sleep(1 * time.Second)

	// Stop server
	close(stop)
}
