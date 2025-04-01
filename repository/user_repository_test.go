package repository

import (
	"math/big"
	"testing"
)

func TestUserRepository_GetUserByUsername_Found(t *testing.T) {
	repo := NewUserRepository()

	user := repo.GetUserByUsername("Alice")

	if user == nil || user.GetUsername() != "Alice" {
		t.Errorf("expected user Alice, got %v", user)
	}
}

func TestUserRepository_GetUserByUsername_NotFound(t *testing.T) {
	repo := NewUserRepository()

	user := repo.GetUserByUsername("Charlie") // Charlie doesn't exist

	if user.GetUsername() != "Default User" {
		t.Errorf("expected Default User, got %s", user.GetUsername())
	}
}

func TestUserRepository_GetUserBalance_Found(t *testing.T) {
	repo := NewUserRepository()

	balance := repo.GetUserBalanceById(1) // Alice has default 0 balance

	expected := big.NewFloat(0.00)
	if balance.Cmp(expected) != 0 {
		t.Errorf("expected balance %.2f, got %.2f", expected, balance)
	}
}

func TestUserRepository_GetUserBalance_NotFound(t *testing.T) {
	repo := NewUserRepository()

	balance := repo.GetUserBalanceById(99) // ID 99 doesn't exist

	expected := big.NewFloat(0.00)
	if balance.Cmp(expected) != 0 {
		t.Errorf("expected balance %.2f for non-existent user, got %.2f", expected, balance)
	}
}
