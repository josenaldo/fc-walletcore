package database

import (
	"database/sql"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

type TransactionDB struct {
	*BaseRepository
}

func NewTransactionDb(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		BaseRepository: &BaseRepository{
			db: db,
		},
	}
}

func NewTransactionDbWithTx(tx *sql.Tx) *TransactionDB {
	return &TransactionDB{
		BaseRepository: &BaseRepository{
			tx: tx,
		},
	}
}

func (repo *TransactionDB) Create(transaction *entity.Transaction) error {
	query := "INSERT INTO transactions (id, created_at, amount, account_from_id, account_to_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := repo.prepareQuery(query)
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
