package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	// Act - When

	// Arrange - Given
	client, err := NewClient("Zé Galinha", "ze@galinha.com")

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, "Zé Galinha", client.Name)
	assert.Equal(t, "ze@galinha.com", client.Email)
}

func TestCreateNewClientWithEmptyName(t *testing.T) {
	// Act - When

	// Arrange - Given
	client, err := NewClient("", "ze@galinha.com")

	// Assert - Then
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Error(t, err, ErrorClientNameIsRequired)
}

func TestCreateNewClientWithEmptyEmail(t *testing.T) {
	// Act - When

	// Arrange - Given
	client, err := NewClient("Zé Galinha", "")

	// Assert - Then
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Error(t, err, ErrorClientEmailIsRequired)
}

func TestUpdateClient(t *testing.T) {
	// Act - When
	client, _ := NewClient("Zé Galinha", "ze@galinha.com")
	updatedBefore := client.UpdatedAt

	// Arrange - Given
	err := client.Update("Zé Galinha da Silva", "ze@galo.com")

	// Assert - Then
	assert.Nil(t, err)
	assert.Equal(t, "Zé Galinha da Silva", client.Name)
	assert.Equal(t, "ze@galo.com", client.Email)
	assert.NotEqual(t, updatedBefore, client.UpdatedAt)
}

func TestUpdateClientWithEmptyName(t *testing.T) {
	// Act - When
	client, _ := NewClient("Zé Galinha", "ze@galinha.com")
	updatedBefore := client.UpdatedAt

	// Arrange - Given
	err := client.Update("", "ze@galinha.com")

	// Assert - Then
	assert.Error(t, err, ErrorClientNameIsRequired)
	assert.Equal(t, "Zé Galinha", client.Name)
	assert.Equal(t, "ze@galinha.com", client.Email)
	assert.Equal(t, updatedBefore, client.UpdatedAt)
}

func TestUpdateClientWithEmptyEmail(t *testing.T) {
	// Act - When
	client, _ := NewClient("Zé Galinha", "ze@galinha.com")
	updatedBefore := client.UpdatedAt

	// Arrange - Given
	err := client.Update("Zé Galinha", "")

	// Assert - Then
	assert.NotNil(t, err)
	assert.Error(t, err, ErrorClientEmailIsRequired)
	assert.Equal(t, "Zé Galinha", client.Name)
	assert.Equal(t, "ze@galinha.com", client.Email)
	assert.Equal(t, updatedBefore, client.UpdatedAt)
}

func TestAddAccountToClient(t *testing.T) {
	// Arrange - Given
	client, _ := NewClient("Zé Galinha", "t@t.com")
	account, _ := NewAccount(client)

	// Act - When
	err := client.AddAccount(account)

	// Assert - Then
	assert.Nil(t, err)
	assert.Len(t, client.Accounts, 1)
	assert.Equal(t, account, client.Accounts[0])
}

func TestAddAccountToClientWithDifferentClient(t *testing.T) {
	// Arrange - Given
	client, _ := NewClient("Zé Galinha", "t@t.com")
	otherClient, _ := NewClient("Zé Pato", "p@p.com")
	account, _ := NewAccount(otherClient)

	// Act - When
	err := client.AddAccount(account)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Error(t, err, ErrorAccountBelongsToAnotherClient)
	assert.Len(t, client.Accounts, 0)
}

func TestAddAccountToClientWithNilAccount(t *testing.T) {
	// Arrange - Given
	client, _ := NewClient("Zé Galinha", "t@t.com")

	// Act - When
	err := client.AddAccount(nil)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Error(t, err, ErrorAccountIsRequired)
	assert.Len(t, client.Accounts, 0)
}
