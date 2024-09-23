package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
    // Start a TCP server to listen for incoming connections from A
    ln, err := net.Listen("tcp", ":8081") // Listening on port 8081
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer ln.Close()
    fmt.Println("System B listening on port 8081")

    for {
        connA, err := ln.Accept() // Accept connection from A
        if err != nil {
            fmt.Println("Error accepting connection from A:", err)
            continue
        }
        go handleConnection(connA)
    }
}

func handleConnection(connA net.Conn) {
    defer connA.Close()

    // Connect to System C
    connC, err := net.Dial("tcp", "localhost:8082") // Assuming C is running on port 8082
    if err != nil {
        fmt.Println("Error connecting to System C:", err)
        return
    }
    defer connC.Close()

    // Forward data from A to C
    _, err = io.Copy(connC, connA) // Copy data from A's connection to C's connection
    if err != nil {
        fmt.Println("Error forwarding file to System C:", err)
        return
    }

    fmt.Println("File forwarded to System C successfully.")
}