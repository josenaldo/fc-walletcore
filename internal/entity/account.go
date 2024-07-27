package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.NewString(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return errors.New("client is required")
	}

	return nil
}

func (a *Account) Credit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	a.Balance += amount
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Account) Debit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if a.Balance < amount {
		return errors.New("insufficient funds")
	}

	a.Balance -= amount
	a.UpdatedAt = time.Now()

	return nil
}
