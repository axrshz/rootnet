package tcp_test

import (
	"net"
	"testing"
	"time"

	"github.com/axrshz/rootnet/pkg/tcp"
)

func TestTCPServer(t *testing.T) {
	go tcp.StartServer("localhost:8080")

	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)	
	}
	defer conn.Close()
	t.Log("Successfully connected to the server")
}