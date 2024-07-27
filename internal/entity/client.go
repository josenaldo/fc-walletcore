package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string
	Email    string
	Accounts []*Account
}

func NewClient(name, email string) (*Client, error) {

	client := &Client{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Name:     name,
		Email:    email,
		Accounts: []*Account{},
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name cannot be empty")
	}

	if c.Email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}

func (c *Client) Update(name, email string) error {
	oldName := c.Name
	oldEmail := c.Email
	oldUpdatedAt := c.UpdatedAt

	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	error := c.Validate()
	if error != nil {
		c.Name = oldName
		c.Email = oldEmail
		c.UpdatedAt = oldUpdatedAt
		return error
	}

	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account == nil {
		return errors.New("account is required")
	}

	if account.Client.ID != c.ID {
		return errors.New("account must belong to the client")
	}

	c.Accounts = append(c.Accounts, account)
	return nil
}
