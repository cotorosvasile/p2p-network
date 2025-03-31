package cli

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"p2p-network/network"
	"p2p-network/service"
	"strconv"
	"strings"
)

// HandleUserInput listens for user's commands and processes them
func HandleUserInput(client *network.Client, userService *service.UserServiceImpl) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case input == "exit":
			fmt.Println("Goodbye.")
			client.Connection.Close()
			return

		case input == "balance":
			balance := userService.ViewBalance(client.Name)
			fmt.Printf("Current balance: %.2f\n", balance)

		case strings.HasPrefix(input, "pay "):
			parts := strings.Split(input, " ")
			if len(parts) != 2 {
				fmt.Println("Usage: pay <amount>")
				continue
			}

			amount, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Println("Invalid amount")
				continue
			}

			/*
				Instead of hardcoding the recipient, I can modify the logic to retrieve a list of connected clients and
				allow the user to choose whom to send money to.
				Steps to Fix:
				1) Store Connected Clients: My Peer struct already maintains a Clients map. We can expose a method
				to return the list of connected clients.
				2) Modify the Pay Command: Instead of always assuming "Alice" pays "Bob" (or vice versa), let users
				pick from available peers.
				3) I was needed to modify the pay command from pay <amount> to pay <recipient> <amount> given that in
				the Terminal Output of the Challenge I saw an example with pay <amount> it is not clear to me what exactly
				is wanted, hardcoding or a dynamic approach.
			*/

			// Get default recipient
			var receiver string
			if client.Name == "Alice" {
				receiver = "Bob"
			} else {
				receiver = "Alice"
			}

			// Process payment
			err = userService.SendMoney(client.Name, big.NewFloat(amount))
			if err != nil {
				fmt.Println("Error sending money:", err)
				return
			}

			fmt.Println("Sent")

			// Send message to notify the receiver
			err = client.SendMessage(fmt.Sprintf("pay %s %.2f", receiver, amount))
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		default:
			fmt.Println("Unknown command. Use 'balance', 'pay <amount>', or 'exit'.")
		}
	}
}
