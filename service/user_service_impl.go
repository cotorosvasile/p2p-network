package service

import (
	"fmt"
	"math/big"
	"p2p-network/repository"
)

type UserServiceImpl struct {
	userRepo *repository.UserRepository
}

func NewUserServiceImpl(userRepo *repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// ViewBalance retrieves the balance of the user by username
func (service *UserServiceImpl) ViewBalance(username string) *big.Float {
	user := service.userRepo.GetUserByUsername(username)
	return user.GetAccountBalance()
}

// SendMoney deducts money from the sender's account
func (service *UserServiceImpl) SendMoney(senderName string, amount *big.Float) error {
	sender := service.userRepo.GetUserByUsername(senderName)
	if sender == nil {
		return fmt.Errorf("sender not found: %s", senderName)
	}

	// Deduct from sender's balance
	currentBalance := sender.GetAccountBalance()
	newBalance := new(big.Float).Sub(currentBalance, amount)
	sender.SetAccountBalance(newBalance)

	return nil
}

// ReceiveMoney adds money to the receiver's account
func (service *UserServiceImpl) ReceiveMoney(receiverName string, amount *big.Float) error {
	receiver := service.userRepo.GetUserByUsername(receiverName)
	if receiver == nil {
		return fmt.Errorf("receiver not found: %s", receiverName)
	}

	// Add to receiver's balance
	currentBalance := receiver.GetAccountBalance()
	newBalance := new(big.Float).Add(currentBalance, amount)
	receiver.SetAccountBalance(newBalance)

	return nil
}
