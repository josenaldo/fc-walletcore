package database

import (
	"database/sql"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

type ClientDb struct {
	Db *sql.DB
}

func NewClientDb(db *sql.DB) *ClientDb {
	return &ClientDb{Db: db}
}

func (c *ClientDb) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.Db.Prepare("SELECT id, created_at, updated_at, name, email FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&client.ID, &client.CreatedAt, &client.UpdatedAt, &client.Name, &client.Email)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func (c *ClientDb) Save(client *entity.Client) error {
	stmt, err := c.Db.Prepare("INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(client.ID, client.CreatedAt, client.UpdatedAt, client.Name, client.Email)
	if err != nil {
		return err
	}

	return nil
}
