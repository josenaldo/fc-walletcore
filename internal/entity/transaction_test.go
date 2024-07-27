package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	john *Client
	jane *Client

	accountJohn *Account
	accountJane *Account
)

func setupTransactionTest() {

	john, _ = NewClient("John Doe", "john@t.com")
	jane, _ = NewClient("Jane Doe", "jane@t.com")

	accountJohn, _ = NewAccount(john)
	accountJane, _ = NewAccount(jane)
}

func TestCreateNewTransaction(t *testing.T) {
	// Act - When
	setupTransactionTest()
	accountJohn.Credit(100)
	accountJane.Credit(50)

	// Arrange - Given
	transaction, err := NewTransaction(accountJohn, accountJane, 25)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.NotEmpty(t, transaction.ID)
	assert.NotNil(t, transaction.CreatedAt)
	assert.Equal(t, accountJohn, transaction.AccountFrom)
	assert.Equal(t, accountJane, transaction.AccountTo)
	assert.Equal(t, 25.0, transaction.Amount)

	assert.Equal(t, 75.0, accountJohn.Balance)
	assert.Equal(t, 75.0, accountJane.Balance)
}

func TestCreateNewTransactionWithEmptyAccountFrom(t *testing.T) {
	// Arrange - Given
	setupTransactionTest()

	// Act - When
	transaction, err := NewTransaction(nil, accountJane, 25)

	// Assert - Then
	assert.Error(t, err, ErrorTransactionAccountFromIsRequired)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithEmptyAccountTo(t *testing.T) {
	// Arrange - Given
	setupTransactionTest()

	// Act - When
	transaction, err := NewTransaction(accountJohn, nil, 25)

	// Assert - Then
	assert.Error(t, err, ErrorTransactionAccountToIsRequired)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithSameAccount(t *testing.T) {
	// Arrange - Given
	setupTransactionTest()

	// Act - When
	transaction, err := NewTransaction(accountJohn, accountJohn, 25)

	// Assert - Then
	assert.Error(t, err, ErrorTansactionAccounFromAndAccountToMustBeDifferent)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithNegativeAmount(t *testing.T) {
	// Arrange - Given
	setupTransactionTest()

	// Act - When
	transaction, err := NewTransaction(accountJohn, accountJane, -25)

	// Assert - Then
	assert.Error(t, err, ErrorTransactionAmountMustBeGreaterThanZero)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithInsufficientFunds(t *testing.T) {
	// Arrange - Given
	setupTransactionTest()
	accountJohn.Credit(100)
	accountJane.Credit(50)

	// Act - When
	transaction, err := NewTransaction(accountJohn, accountJane, 200)

	// Assert - Then
	assert.Error(t, err, ErrorInsufficientFunds)
	assert.Nil(t, transaction)
}
