package repository

import (
	"math/big"
	"p2p-network/entity"
)

type UserRepository struct {
	users map[int]*entity.User // Using a map for quick lookup
}

func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		users: make(map[int]*entity.User),
	}

	repo.users[1] = entity.NewUser(1, "Alice", big.NewFloat(0.00))
	repo.users[2] = entity.NewUser(2, "Bob", big.NewFloat(0.00))

	return repo
}

func (repository *UserRepository) GetUserByUsername(username string) *entity.User {
	for _, user := range repository.users {
		if user.GetUsername() == username {
			return user
		}
	}

	return entity.NewUser(0, "Default User", big.NewFloat(0.00))
}

func (repository *UserRepository) GetUserById(id int) *entity.User {
	for _, user := range repository.users {
		if user.GetID() == id {
			return user
		}
	}

	return entity.NewUser(0, "Default User", big.NewFloat(0.00))
}

func (repository *UserRepository) GetUserBalance(userID int) *big.Float {
	user, exists := repository.users[userID]
	if !exists {
		return big.NewFloat(0.00)
	}

	return user.GetAccountBalance()
}
