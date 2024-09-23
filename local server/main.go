package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
    // Connect to System B
    conn, err := net.Dial("tcp", "localhost:8081") // Assuming B is running on port 8081
    if err != nil {
        fmt.Println("Error connecting to System B:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Open the file to send
    file, err := os.Open("file_to_send.txt") // Ensure this file exists
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Send the file content to System B
    _, err = io.Copy(conn, file)
    if err != nil {
        fmt.Println("Error sending file:", err)
        return
    }

    fmt.Println("File sent to System B successfully.")
}