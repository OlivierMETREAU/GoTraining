package tcp

import (
	"net"
)

// Client represents a connected user.
// The server only needs to know the nickname and the outgoing channel.
// The connection itself will be handled in client.go later.
type Client struct {
	Nick string
	Conn net.Conn
	Out  chan string
}

// Server holds the global state of the chat server.
type Server struct {
	Clients   map[net.Addr]*Client
	Join      chan *Client
	Leave     chan *Client
	Broadcast chan string
}
