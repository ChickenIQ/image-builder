package main

import (
    "log"
    "net"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT is not set")
    }

    l, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if err == nil {
            conn.Close()
        }
    }
}