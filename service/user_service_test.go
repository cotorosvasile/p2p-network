package service

import (
	"math/big"
	"p2p-network/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock implementation of UserRepository
type MockUserRepository struct {
	users map[string]*entity.User
}

func (m *MockUserRepository) GetUserByUsername(username string) *entity.User {
	if user, exists := m.users[username]; exists {
		return user
	}
	return nil // Simulating user not found
}

func newMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: map[string]*entity.User{
			"Alice": entity.NewUser(1, "Alice", big.NewFloat(100.00)),
			"Bob":   entity.NewUser(2, "Bob", big.NewFloat(50.00)),
		},
	}
}

func TestViewBalance(t *testing.T) {
	mockRepo := newMockUserRepository()
	service := NewUserServiceImpl(mockRepo)

	balance := service.ViewBalance("Alice")
	expectedBalance := big.NewFloat(100.00)

	assert.Equal(t, expectedBalance, balance)
}

func TestSendMoney_Success(t *testing.T) {
	mockRepo := newMockUserRepository()
	service := NewUserServiceImpl(mockRepo)

	err := service.SendMoney("Alice", big.NewFloat(30.00))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedBalance := big.NewFloat(70.00) // Alice had 100, sent 30
	actualBalance := service.ViewBalance("Alice")

	assert.Equal(t, expectedBalance, actualBalance)
}

func TestSendMoney_UserNotFound(t *testing.T) {
	mockRepo := newMockUserRepository()
	service := NewUserServiceImpl(mockRepo)

	err := service.SendMoney("Charlie", big.NewFloat(10.00)) // Charlie doesn't exist

	assert.ErrorContains(t, err, "sender not found", "Expected error to contain 'receiver not found'")
}

func TestReceiveMoney_Success(t *testing.T) {
	mockRepo := newMockUserRepository()
	service := NewUserServiceImpl(mockRepo)

	err := service.ReceiveMoney("Bob", big.NewFloat(20.00))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedBalance := big.NewFloat(70.00) // Bob had 50, received 20
	actualBalance := service.ViewBalance("Bob")

	assert.Equal(t, expectedBalance, actualBalance)
}

func TestReceiveMoney_UserNotFound(t *testing.T) {
	mockRepo := newMockUserRepository()
	service := NewUserServiceImpl(mockRepo)

	err := service.ReceiveMoney("Charlie", big.NewFloat(10.00)) // Charlie doesn't exist

	assert.ErrorContains(t, err, "receiver not found", "Expected error to contain 'receiver not found'")
}
