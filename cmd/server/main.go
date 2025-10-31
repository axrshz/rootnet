package main

import (
	"fmt"
	"io"
	"net"

	"github.com/axrshz/rootnet/pkg/http"
	"github.com/axrshz/rootnet/pkg/http/status"
)

func main() {
	router := http.NewRouter()

	// Add routes
	router.AddRoute("GET", "/", handleRoot)
	router.AddRoute("GET", "/hello", handleHello)
	router.AddRoute("POST", "/echo", handleEcho)

	startServer(":8080", router)
}

func startServer(address string, router *http.Router) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn, router)
	}
}

func handleConnection(conn net.Conn, router *http.Router) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading from connection: %v\n", err)
		return
	}

	request, err := http.ParseRequest(buffer[:n])
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		sendErrorResponse(conn, status.BadRequest)
		return
	}

	response := router.HandleRequest(request)
	_, err = conn.Write(http.FormatResponse(response))
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}

func sendErrorResponse(conn net.Conn, statusCode int) {
	response := http.NewResponse()
	response.StatusCode = statusCode
	response.StatusText = http.StatusText(statusCode)
	response.SetBody([]byte(http.StatusText(statusCode)))

	_, err := conn.Write(http.FormatResponse(response))
	if err != nil {
		fmt.Printf("Error writing error response: %v\n", err)
	}
}

// Handler functions
func handleRoot(req *http.Request) *http.Response {
	resp := http.NewResponse()
	resp.StatusCode = 200
	resp.StatusText = "OK"
	resp.SetBody([]byte("welcome to rootnet!"))
	return resp
}

func handleHello(req *http.Request) *http.Response {
	resp := http.NewResponse()
	resp.StatusCode = 200
	resp.StatusText = "OK"
	resp.SetBody([]byte("hello, rootnet!"))
	return resp
}

func handleEcho(req *http.Request) *http.Response {
	resp := http.NewResponse()
	resp.StatusCode = 200
	resp.StatusText = "OK"
	resp.SetHeader("Content-Type", "text/plain")
	resp.SetHeader("Content-Length", fmt.Sprintf("%d", len(req.Body)))
	resp.Body = req.Body
	return resp
}