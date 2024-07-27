package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Zé Galinha", "ze@galinha.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, "Zé Galinha", client.Name)
	assert.Equal(t, "ze@galinha.com", client.Email)
}

func TestCreateNewClientWithEmptyName(t *testing.T) {
	client, err := NewClient("", "ze@galinha.com")
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "name cannot be empty", err.Error())
}

func TestCreateNewClientWithEmptyEmail(t *testing.T) {
	client, err := NewClient("Zé Galinha", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "email cannot be empty", err.Error())
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Zé Galinha", "ze@galinha.com")
	updatedBefore := client.UpdatedAt

	err := client.Update("Zé Galinha da Silva", "ze@galo.com")

	assert.Nil(t, err)
	assert.Equal(t, "Zé Galinha da Silva", client.Name)
	assert.Equal(t, "ze@galo.com", client.Email)
	assert.NotEqual(t, updatedBefore, client.UpdatedAt)
}

func TestUpdateClientWithEmptyName(t *testing.T) {
	client, _ := NewClient("Zé Galinha", "ze@galinha.com")
	updatedBefore := client.UpdatedAt

	err := client.Update("", "ze@galinha.com")

	assert.Error(t, err, "name cannot be empty")
	assert.Equal(t, "Zé Galinha", client.Name)
	assert.Equal(t, "ze@galinha.com", client.Email)
	assert.Equal(t, updatedBefore, client.UpdatedAt)
}

func TestUpdateClientWithEmptyEmail(t *testing.T) {
	client, _ := NewClient("Zé Galinha", "ze@galinha.com")
	updatedBefore := client.UpdatedAt

	err := client.Update("Zé Galinha", "")

	assert.NotNil(t, err)
	assert.Equal(t, "email cannot be empty", err.Error())
	assert.Equal(t, "Zé Galinha", client.Name)
	assert.Equal(t, "ze@galinha.com", client.Email)
	assert.Equal(t, updatedBefore, client.UpdatedAt)
}
