package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "secure-p2p-platform/traffic"
    "secure-p2p-platform/encryption"
)

func main() {
    cert, key := "server.crt", "server.key"
    listener := encryption.StartSecureServer(":8443", cert, key)
    defer listener.Close()

    fmt.Println("Secure P2P server listening on port 8443...")
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Connection error: %s", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Printf("Client connected: %s\n", conn.RemoteAddr())
    traffic.LogTraffic(fmt.Sprintf("New connection from %s", conn.RemoteAddr()))
    conn.Close()
}
