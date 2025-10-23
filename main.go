package main

import (
	"fmt"
	"net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer ln.Close()

    fmt.Println("listening on :8080")
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            continue
        }
        go handle(conn)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("read error:", err)
        return
    }
    fmt.Println("received:", string(buf[:n]))
    conn.Write([]byte("hello tcp\n"))
}
