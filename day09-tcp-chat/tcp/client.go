package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// HandleClient attaches a TCP connection to the server.
// It creates a Client, registers it, and starts reader/writer goroutines.
func HandleClient(s *Server, conn net.Conn) {
	client := &Client{
		Nick: conn.RemoteAddr().String(),
		Conn: conn,
		Out:  make(chan string, 16),
	}

	// Register client in the server
	s.Join <- client

	// Start writer goroutine (server → TCP)
	go clientWriter(client)

	// Start reader loop (TCP → server)
	clientReader(s, client)
}

// --- Writer goroutine ---
// Sends messages from client.Out to the TCP connection.
func clientWriter(c *Client) {
	writer := bufio.NewWriter(c.Conn)

	for msg := range c.Out {
		_, err := writer.WriteString(msg + "\n")
		if err != nil {
			break
		}
		writer.Flush()
	}

	// When Out is closed, close the TCP connection
	c.Conn.Close()
}

// --- Reader loop ---
// Reads lines from TCP and forwards them to the server.
func clientReader(s *Server, c *Client) {
	scanner := bufio.NewScanner(c.Conn)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Commands
		if strings.HasPrefix(line, "/") {
			handleCommand(s, c, line)
			continue
		}

		// Normal chat message
		s.Broadcast <- fmt.Sprintf("[%s] %s", c.Nick, line)
	}

	// On disconnect or error
	s.Leave <- c
}

// --- Command handling ---
func handleCommand(s *Server, c *Client, line string) {
	switch {
	case strings.HasPrefix(line, "/nick "):
		newNick := strings.TrimSpace(strings.TrimPrefix(line, "/nick "))
		if newNick == "" {
			c.Out <- "Usage: /nick <name>"
			return
		}
		old := c.Nick
		c.Nick = newNick
		s.Broadcast <- fmt.Sprintf("* %s is now known as %s", old, newNick)

	case line == "/who":
		c.Out <- "--- Connected users ---"
		for _, cl := range s.Clients {
			c.Out <- " - " + cl.Nick
		}
		c.Out <- "-----------------------"

	case line == "/quit":
		c.Out <- "Bye!"
		s.Leave <- c
		c.Conn.Close()

	default:
		c.Out <- "Unknown command"
	}
}
