package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name  string
	email string
)

func setupClientTest() {
	name = "Zé Galinha"
	email = "ze@galinha.com"
}

func TestCreateNewClient(t *testing.T) {
	// Arrange - Given
	setupClientTest()

	// Act - When
	client, err := NewClient(name, email)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, name, client.Name)
	assert.Equal(t, email, client.Email)
}

func TestCreateNewClientWithEmptyName(t *testing.T) {
	// Arrange - Given
	setupClientTest()

	// Act - When
	client, err := NewClient("", email)

	// Assert - Then
	assert.EqualError(t, err, ErrorClientNameIsRequired.Error())
	assert.Nil(t, client)
}

func TestCreateNewClientWithEmptyEmail(t *testing.T) {
	// Arrange - Given
	setupClientTest()

	// Act - When
	client, err := NewClient(name, "")

	// Assert - Then
	assert.EqualError(t, err, ErrorClientEmailIsRequired.Error())
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	// Arrange - Given
	setupClientTest()
	client, _ := NewClient(name, email)
	updatedBefore := client.UpdatedAt

	// Act - When
	err := client.Update("Zé Galinha da Silva", "ze@galo.com")

	// Assert - Then
	assert.Nil(t, err)
	assert.Equal(t, "Zé Galinha da Silva", client.Name)
	assert.Equal(t, "ze@galo.com", client.Email)
	assert.NotEqual(t, updatedBefore, client.UpdatedAt)
}

func TestUpdateClientWithEmptyName(t *testing.T) {
	// Arrange - Given
	setupClientTest()
	client, _ := NewClient(name, email)
	updatedBefore := client.UpdatedAt

	// Act - When
	err := client.Update("", email)

	// Assert - Then
	assert.EqualError(t, err, ErrorClientNameIsRequired.Error())
	assert.Equal(t, name, client.Name)
	assert.Equal(t, email, client.Email)
	assert.Equal(t, updatedBefore, client.UpdatedAt)
}

func TestUpdateClientWithEmptyEmail(t *testing.T) {
	// Arrange - Given
	setupClientTest()
	client, _ := NewClient(name, email)
	updatedBefore := client.UpdatedAt

	// Act - When
	err := client.Update(name, "")

	// Assert - Then
	assert.EqualError(t, err, ErrorClientEmailIsRequired.Error())
	assert.Equal(t, name, client.Name)
	assert.Equal(t, email, client.Email)
	assert.Equal(t, updatedBefore, client.UpdatedAt)
}

func TestAddAccountToClient(t *testing.T) {
	// Arrange - Given
	setupClientTest()
	client, _ := NewClient(name, "t@t.com")

	// Act - When
	account, err := NewAccount(client)

	// Assert - Then
	assert.Nil(t, err)
	assert.Len(t, client.Accounts, 1)
	assert.Equal(t, account, client.Accounts[0])
}

func TestAddAccountToClientWithDifferentClient(t *testing.T) {
	// Arrange - Given
	setupClientTest()
	client, _ := NewClient(name, "t@t.com")
	otherClient, _ := NewClient("Zé Pato", "p@p.com")
	account, _ := NewAccount(otherClient)

	// Act - When
	err := client.AddAccount(account)

	// Assert - Then
	assert.EqualError(t, err, ErrorAccountBelongsToAnotherClient.Error())
	assert.Len(t, client.Accounts, 0)
}

func TestAddAccountToClientWithNilAccount(t *testing.T) {
	// Arrange - Given
	setupClientTest()
	client, _ := NewClient(name, email)

	// Act - When
	err := client.AddAccount(nil)

	// Assert - Then
	assert.EqualError(t, err, ErrorAccountIsRequired.Error())
	assert.Len(t, client.Accounts, 0)
}
