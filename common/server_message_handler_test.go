package common

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"p2p-network/entity"
	"p2p-network/service"
	"testing"
)

// MockUserRepository mocks the UserRepository interface
type MockUserRepository struct {
	mock.Mock
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

func TestHandlePaymentNotification_ValidPayment(t *testing.T) {
	mockRepo := newMockUserRepository()
	mockRepo.On("GetUserByUsername", "Alice").Return(mockRepo.users["Alice"])

	userService := service.NewUserServiceImpl(mockRepo)

	HandlePaymentNotification("You were paid 10.00 !", "Alice", userService)

	actualBalance := userService.ViewBalance("Alice")

	// Check if Alice's balance is updated
	expectedBalance := big.NewFloat(110.00)
	assert.Equal(t, expectedBalance, actualBalance, "Balance should be updated correctly")
}

func TestHandlePaymentNotification_NonNumericAmount(t *testing.T) {
	mockRepo := newMockUserRepository()
	userService := service.NewUserServiceImpl(mockRepo)

	HandlePaymentNotification("You were paid abc!", "Alice", userService)

	actualBalance := userService.ViewBalance("Alice")

	// Ensure Alice's balance remains unchanged
	expectedBalance := big.NewFloat(100.00)
	assert.Equal(t, expectedBalance, actualBalance, "Balance should not change for invalid message")
}

func TestHandlePaymentNotification_IrrelevantMessage(t *testing.T) {
	mockRepo := newMockUserRepository()
	userService := service.NewUserServiceImpl(mockRepo)

	HandlePaymentNotification("Hello Alice, how are you?", "Alice", userService)

	actualBalance := userService.ViewBalance("Alice")

	// Ensure Alice's balance remains unchanged
	expectedBalance := big.NewFloat(100.00)
	assert.Equal(t, expectedBalance, actualBalance, "Balance should not change for unrelated message")
}
