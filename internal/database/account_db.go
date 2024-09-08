package database

import (
	"database/sql"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDb(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (db *AccountDB) Get(id string) (*entity.Account, error) {

	stmt, err := db.DB.Prepare(`
	SELECT 
		a.id, 
		a.created_at, 
		a.updated_at, 		
		a.balance, 
		c.id, 
		c.created_at, 
		c.updated_at, 
		c.name, 
		c.email
	FROM 
		accounts a INNER JOIN  clients c ON a.client_id = c.id
	WHERE 
		a.id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var account entity.Account
	var client entity.Client
	account.Client = &client

	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.Balance,
		&client.ID,
		&client.CreatedAt,
		&client.UpdatedAt,
		&client.Name,
		&client.Email)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (db *AccountDB) Save(account *entity.Account) error {
	insertQuery := `
		INSERT 
			INTO accounts (id, created_at, updated_at, balance, client_id) 
			VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.CreatedAt, account.UpdatedAt, account.Balance, account.Client.ID)
	if err != nil {
		return err
	}

	return nil
}
