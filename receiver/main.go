package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
    // Start a TCP server to listen for incoming connections from B
    ln, err := net.Listen("tcp", ":8082") // Listening on port 8082
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer ln.Close()
    fmt.Println("System C listening on port 8082")

    for {
        conn, err := ln.Accept() // Accept connection from B
        if err != nil {
            fmt.Println("Error accepting connection from B:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Create a file to save the received data
    file, err := os.Create("received_file.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Copy data from the connection to the file
    _, err = io.Copy(file, conn)
    if err != nil {
        fmt.Println("Error receiving file data:", err)
        return
    }

    fmt.Println("File received from System B successfully.")
}