package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	client *Client
)

func TestMain(m *testing.M) {
	// Setup
	client, _ = NewClient("Zé Galinha", "t@t.com")

	// Run tests
	m.Run()

	// Teardown
}

func TestCreateNewAccount(t *testing.T) {
	// Arrange - Given

	// Act - When
	account, err := NewAccount(client)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.NotEmpty(t, account.ID)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, 0.0, account.Balance)
	assert.NotNil(t, account.CreatedAt)
	assert.NotNil(t, account.UpdatedAt)
}

func TestCreateNewAccountWithEmptyClient(t *testing.T) {
	// Arrange - Given

	// Act - When
	account, err := NewAccount(nil)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.Equal(t, "client is required", err.Error())
}

func TestCreditAccount(t *testing.T) {
	// Arrange - Given
	account, _ := NewAccount(client)

	// Act - When
	err := account.Credit(100)

	// Assert - Then
	assert.Nil(t, err)
	assert.Equal(t, 100.0, account.Balance)
}

func TestCreditAccountWithNegativeAmount(t *testing.T) {
	// Arrange - Given
	account, _ := NewAccount(client)

	// Act - When
	err := account.Credit(-100)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than 0", err.Error())
}

func TestDebitAccount(t *testing.T) {
	// Arrange - Given
	account, _ := NewAccount(client)
	account.Credit(100)

	// Act - When
	err := account.Debit(50)

	// Assert - Then
	assert.Nil(t, err)
	assert.Equal(t, 50.0, account.Balance)
}

func TestDebitAccountWithNegativeAmount(t *testing.T) {
	// Arrange - Given
	account, _ := NewAccount(client)

	// Act - When
	err := account.Debit(-100)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than 0", err.Error())
}

func TestDebitAccountWithInsufficientFunds(t *testing.T) {
	// Arrange - Given
	account, _ := NewAccount(client)
	account.Credit(100)

	// Act - When
	err := account.Debit(150)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Equal(t, "insufficient funds", err.Error())
}

func TestDebitAccountWithZeroAmount(t *testing.T) {
	// Arrange - Given
	account, _ := NewAccount(client)
	account.Credit(100)

	// Act - When
	err := account.Debit(0)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than 0", err.Error())
}
