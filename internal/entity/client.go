package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name, email string) (*Client, error) {

	client := &Client{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
