package common

import (
	"fmt"
	"math/big"
	"p2p-network/service"
	"strconv"
	"strings"
)

// HandlePaymentNotification checks if the response is a payment notification and updates balance
func HandlePaymentNotification(response, receiver string, userService *service.UserServiceImpl) {
	trimmedResponse := strings.TrimSpace(response)
	if strings.Contains(trimmedResponse, "paid ") {
		parts := strings.Split(trimmedResponse, " ")
		if len(parts) == 5 { // Expected format: "You were paid <amount>!"
			amountStr := parts[3]

			// Convert amount to float
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid payment amount received.")
				return
			}

			// Invoke ReceiveMoney
			err = userService.ReceiveMoney(receiver, big.NewFloat(amount))
			if err != nil {
				fmt.Println("Error updating balance.")
				return
			}
		}
	}
}
