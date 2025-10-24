package tcp_test

import (
	"net"
	"testing"
	"time"

	"github.com/axrshz/rootnet/pkg/tcp"
)

func TestTCPServer(t *testing.T) {
	// Start the server
	go tcp.StartServer("localhost:8080")

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Try to connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	t.Log("Successfully connected to the server")
}
