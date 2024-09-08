package database

import (
	"database/sql"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDb(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (db *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := db.DB.Prepare("INSERT INTO transactions (id, created_at, amount, account_from_id, account_to_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(transaction.ID, transaction.CreatedAt, transaction.Amount, transaction.AccountFrom.ID, transaction.AccountTo.ID)
	if err != nil {
		return err
	}

	return nil
}
