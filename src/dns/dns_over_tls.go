package dns

import (
    "crypto/tls"
    "log"
    "net"
)

func (ds *DNSServer) StartDNSoverTLS() error {
    cert, err := tls.LoadX509KeyPair("path/to/cert.pem", "path/to/key.pem")
    if err != nil {
        return err
    }

    config := &tls.Config{Certificates: []tls.Certificate{cert}}
    ln, err := tls.Listen("tcp", ":853", config)
    if err != nil {
        return err
    }
    defer ln.Close()

    log.Println("DNS over TLS started on :853")
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        go ds.handleTLSConnection(conn)
    }
}

func (ds *DNSServer) handleTLSConnection(conn net.Conn) {
    defer conn.Close()
    // Логика обработки DNS-запросов через TLS
}

