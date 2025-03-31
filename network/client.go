package network

import (
	"bufio"
	"fmt"
	"net"
	"p2p-network/common"
	"p2p-network/service"
)

// Client represents a client connection to a peer
type Client struct {
	ServerAddress string
	Name          string
	Connection    net.Conn
}

// NewClient creates a new client
func NewClient(serverAddress, name string) *Client {
	return &Client{
		ServerAddress: serverAddress,
		Name:          name,
	}
}

// Connect establishes a connection to the server
func (c *Client) Connect(userService *service.UserServiceImpl) error {
	conn, err := net.Dial(ServerType, c.ServerAddress)
	if err != nil {
		return err
	}

	c.Connection = conn
	fmt.Println("Welcome to your peering relationship!")

	// Send client name to server for identification
	_, err = c.Connection.Write([]byte(c.Name + "\n"))
	if err != nil {
		return err
	}

	// Start a goroutine to read responses from the server
	go c.ReadResponses(userService)

	return nil
}

// SendMessage sends a message to the server
func (c *Client) SendMessage(message string) error {
	if c.Connection == nil {
		return fmt.Errorf("not connected to server")
	}

	_, err := c.Connection.Write([]byte(message + "\n"))
	return err
}

// ReadResponses reads and displays messages from the server
func (c *Client) ReadResponses(userService *service.UserServiceImpl) {
	reader := bufio.NewReader(c.Connection)

	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		fmt.Printf("%s", response)

		// Delegate payment processing to a separate handler
		common.HandlePaymentNotification(response, c.Name, userService)
	}
}
