package main

import (
	"p2p-network/cli"
	"p2p-network/network"
	"p2p-network/repository"
	"p2p-network/service"
)

const (
	UserBob = "Bob"
)

func main() {

	// Initialize user repository and service
	userRepo := repository.NewUserRepository()
	userService := service.NewUserServiceImpl(userRepo)

	bob := network.NewClient(":8000", UserBob)
	err := bob.Connect(userService)
	if err != nil {
		panic(err)
	}

	// Handle input from CLI
	cli.HandleUserInput(bob, userService)
}
