package entity

import "math/big"

type User struct {
	id             int
	userName       string
	accountBalance *big.Float
}

func NewUser(userID int, username string, initialBalance *big.Float) *User {
	return &User{
		id:             userID,
		userName:       username,
		accountBalance: initialBalance,
	}
}

func (u *User) GetID() int {
	return u.id
}

func (u *User) GetUsername() string {
	return u.userName
}

func (u *User) GetAccountBalance() *big.Float {
	return u.accountBalance
}

func (u *User) SetAccountBalance(newBalance *big.Float) {
	u.accountBalance = newBalance
}
