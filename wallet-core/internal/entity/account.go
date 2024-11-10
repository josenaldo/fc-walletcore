package entity

import (
	"errors"
	"time"
)

var (
	ErrorClientIsRequired            = errors.New("client is required")
	ErrorAmountMustBeGreaterThanZero = errors.New("amount must be greater than 0")
	ErrorInsufficientFunds           = errors.New("insufficient funds")
)

type Account struct {
	ID        EntityID
	CreatedAt time.Time
	UpdatedAt time.Time

	Client  *Client
	Balance float64
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        NewEntityID(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	client.AddAccount(account)

	return account, nil
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return ErrorClientIsRequired
	}

	return nil
}

func (a *Account) Credit(amount float64) error {
	if amount <= 0 {
		return ErrorAmountMustBeGreaterThanZero
	}

	a.Balance += amount
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Account) Debit(amount float64) error {
	if amount <= 0 {
		return ErrorAmountMustBeGreaterThanZero
	}

	if a.Balance < amount {
		return ErrorInsufficientFunds
	}

	a.Balance -= amount
	a.UpdatedAt = time.Now()

	return nil
}
