package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	ServerType = "tcp"
)

type Peer struct {
	Address string
	Clients map[string]net.Conn
}

// NewPeer creates a new Peer with a given address
func NewPeer(address string) *Peer {
	return &Peer{
		Address: address,
		Clients: make(map[string]net.Conn),
	}
}

// StartServer starts listening for incoming connections
func (p *Peer) StartServer() {
	ln, err := net.Listen(ServerType, p.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started. Listening on", p.Address)

	// Accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		clientAddr := conn.RemoteAddr().String()
		p.Clients[clientAddr] = conn
		fmt.Printf("New client connected: %s\n", clientAddr)

		// Handle the connection in a new goroutine
		go p.HandleConnection(conn)
	}
}

// HandleConnection processes incoming connections
func (p *Peer) HandleConnection(conn net.Conn) {
	defer func() {
		clientAddr := conn.RemoteAddr().String()
		delete(p.Clients, clientAddr)
		fmt.Printf("Client disconnected: %s\n", clientAddr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	// Read the first message as the client name
	clientName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading client name:", err)
		return
	}
	clientName = strings.TrimSpace(clientName)

	// Store client using name instead of IP
	p.Clients[clientName] = conn

	// Handle incoming messages
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		trimmedMessage := strings.TrimSpace(message)
		fmt.Printf("Message from %s: %s\n", clientName, trimmedMessage)

		// Handle payment message correctly
		if strings.HasPrefix(trimmedMessage, "pay ") {
			parts := strings.Split(trimmedMessage, " ")
			if len(parts) == 3 {
				recipient := parts[1]
				amount := parts[2]

				// Send payment notification to the correct receiver
				receiver, exists := p.Clients[recipient]
				if exists {
					notification := fmt.Sprintf("You were paid %s !\n", amount)
					_, err := receiver.Write([]byte(notification))
					if err != nil {
						fmt.Printf("Error sending notification to %s: %s\n", recipient, err)
					} else {
						fmt.Printf("Payment notification sent to %s\n", recipient)
					}
				} else {
					fmt.Printf("Recipient %s not found\n", recipient)
				}
			}
		}
	}
}
