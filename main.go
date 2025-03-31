package main

import (
	"p2p-network/network"
)

const (
	ServerPort = ":8000"
)

func main() {
	server := network.NewPeer(ServerPort)
	server.StartServer()
}
