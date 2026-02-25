package tcp

import "net"

type ServerOption func(*Server)

func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		Clients:   make(map[net.Addr]*Client),
		Join:      make(chan *Client),
		Leave:     make(chan *Client),
		Broadcast: make(chan string),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Run is the main event loop.
// It is intentionally pure: no TCP, no goroutines per client.
// This makes it testable.
func (s *Server) Run(stop <-chan struct{}) {
	for {
		select {
		case c := <-s.Join:
			s.addClient(c)

		case c := <-s.Leave:
			s.removeClient(c)

		case msg := <-s.Broadcast:
			s.broadcastMessage(msg)

		case <-stop:
			return
		}
	}
}

// --- Internal helpers ---

func (s *Server) addClient(c *Client) {
	s.Clients[c.Conn.RemoteAddr()] = c
}

func (s *Server) removeClient(c *Client) {
	if _, ok := s.Clients[c.Conn.RemoteAddr()]; ok {
		delete(s.Clients, c.Conn.RemoteAddr())
		close(c.Out)
	}
}

func (s *Server) broadcastMessage(msg string) {
	for _, c := range s.Clients {
		select {
		case c.Out <- msg:
		default:
			// Drop message if client is slow
		}
	}
}
