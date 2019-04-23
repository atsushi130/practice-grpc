package main

import (
    "fmt"
    "net"
    "google.golang.org/grpc"
)

func main() {
    listenPort, err := net.Listen("tcp", ":19003")
    if err != nil {
        fmt.Printf("エラーだよ")
    }
    server := grpc.NewServer()
    server.Serve(listenPort)
}
