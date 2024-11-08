package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorTransactionAccountFromIsRequired                = errors.New("transaction account from is required")
	ErrorTransactionAccountToIsRequired                  = errors.New("transaction account to is required")
	ErrorTransactionAmountMustBeGreaterThanZero          = errors.New("transaction amount must be greater than zero")
	ErrorTansactionAccounFromAndAccountToMustBeDifferent = errors.New("transaction account from and account to must be different")
)

type Transaction struct {
	ID        string
	CreatedAt time.Time

	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),

		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return ErrorTransactionAccountFromIsRequired
	}

	if t.AccountTo == nil {
		return ErrorTransactionAccountToIsRequired
	}

	if t.AccountFrom.ID == t.AccountTo.ID {
		return ErrorTansactionAccounFromAndAccountToMustBeDifferent
	}

	if t.Amount <= 0 {
		return ErrorTransactionAmountMustBeGreaterThanZero
	}

	return nil
}

func (t *Transaction) Commit() error {
	err := t.AccountFrom.Debit(t.Amount)
	if err != nil {
		return err
	}

	err = t.AccountTo.Credit(t.Amount)
	if err != nil {
		return err
	}

	return nil
}
