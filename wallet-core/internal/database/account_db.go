package database

import (
	"database/sql"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
)

type AccountDB struct {
	*BaseRepository
}

func NewAccountDb(db *sql.DB) *AccountDB {
	return &AccountDB{
		BaseRepository: &BaseRepository{
			db: db,
		},
	}
}

func NewAccountDbWithTx(tx *sql.Tx) *AccountDB {
	return &AccountDB{
		BaseRepository: &BaseRepository{
			tx: tx,
		},
	}
}

func (repo *AccountDB) Get(id entity.EntityID) (*entity.Account, error) {
	query := `
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
			a.id = ?`

	stmt, err := repo.prepareQuery(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var account entity.Account
	var client entity.Client
	account.Client = &client

	row := stmt.QueryRow(id.String())
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
		if err == sql.ErrNoRows {
			return nil, gateway.ErrorAccountNotFound
		}

		return nil, err
	}

	return &account, nil
}

func (repo *AccountDB) Save(account *entity.Account) error {
	query := `
		INSERT 
			INTO accounts (id, created_at, updated_at, balance, client_id) 
			VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := repo.prepareQuery(query)
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

func (repo *AccountDB) UpdateBalance(account *entity.Account) error {
	query := `
		UPDATE accounts 
			SET balance = ?, updated_at = ? 
			WHERE id = ?`

	stmt, err := repo.prepareQuery(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.Balance, account.UpdatedAt, account.ID)
	if err != nil {
		return err
	}

	return nil
}
