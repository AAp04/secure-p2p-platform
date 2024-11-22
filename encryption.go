package encryption

import (
    "crypto/tls"
    "log"
    "net"
)

func StartSecureServer(port, cert, key string) net.Listener {
    tlsCert, err := tls.LoadX509KeyPair(cert, key)
    if err != nil {
        log.Fatalf("Failed to load TLS keys: %s", err)
    }

    config := &tls.Config{Certificates: []tls.Certificate{tlsCert}}
    listener, err := tls.Listen("tcp", port, config)
    if err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
    return listener
}
