package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorClientNameIsRequired          = errors.New("client name is required")
	ErrorClientEmailIsRequired         = errors.New("client email is required")	
	ErrorAccountIsRequired             = errors.New("account is required")
	ErrorAccountBelongsToAnotherClient = errors.New("account must belong to the client")
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
		return ErrorClientNameIsRequired
	}

	if c.Email == "" {
		return ErrorClientEmailIsRequired
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
		return ErrorAccountIsRequired
	}

	if account.Client.ID != c.ID {
		return ErrorAccountBelongsToAnotherClient
	}

	c.Accounts = append(c.Accounts, account)
	return nil
}
