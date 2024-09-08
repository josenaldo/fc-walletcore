package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db            *sql.DB
	TransactionDb *TransactionDB

	clientFrom  *entity.Client
	accountFrom *entity.Account

	clientTo  *entity.Client
	accountTo *entity.Account
}

func (s *TransactionDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	db.Exec(`
	CREATE TABLE clients (
		id VARCHAR(255) PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		name VARCHAR(255),
		email VARCHAR(255)
		)`)

	db.Exec(`
	CREATE TABLE accounts (
		id VARCHAR(255) PRIMARY KEY, 
		created_at DATETIME, 
		updated_at DATETIME, 
		balance DECIMAL, 
		client_id VARCHAR(255)
		)`)

	db.Exec(`
	CREATE TABLE transactions (
		id VARCHAR(255) PRIMARY KEY, 
		created_at DATETIME, 
		amount DECIMAL, 
		account_from_id VARCHAR(255), 
		account_to_id VARCHAR(255)
		)`)

	s.TransactionDb = NewTransactionDb(db)

	s.clientFrom, _ = entity.NewClient("Zé Galinha", "ze@galinha.com")
	s.db.Exec("INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)",
		s.clientFrom.ID, s.clientFrom.CreatedAt, s.clientFrom.UpdatedAt, s.clientFrom.Name, s.clientFrom.Email)

	s.clientTo, _ = entity.NewClient("Maria Galinha", "maria@galinha.com")
	s.db.Exec("INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)",
		s.clientTo.ID, s.clientTo.CreatedAt, s.clientTo.UpdatedAt, s.clientTo.Name, s.clientTo.Email)

	s.accountFrom, _ = entity.NewAccount(s.clientFrom)
	s.accountFrom.Credit(1000)
	s.db.Exec("INSERT INTO accounts (id, created_at, updated_at, balance, client_id) VALUES (?, ?, ?, ?, ?)",
		s.accountFrom.ID, s.accountFrom.CreatedAt, s.accountFrom.UpdatedAt, s.accountFrom.Balance, s.accountFrom.Client.ID)

	s.accountTo, _ = entity.NewAccount(s.clientTo)
	s.accountTo.Credit(1000)
	s.db.Exec("INSERT INTO accounts (id, created_at, updated_at, balance, client_id) VALUES (?, ?, ?, ?, ?)",
		s.accountTo.ID, s.accountTo.CreatedAt, s.accountTo.UpdatedAt, s.accountTo.Balance, s.accountTo.Client.ID)

}

func (s *TransactionDbTestSuite) TearDownSuite() {
	defer s.db.Close()

	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestTransactionDbTestSuit(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) TestCreate() {
	// Arrange - Given
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 200)
	s.Nil(err)

	// Act - When
	err = s.TransactionDb.Create(transaction)

	// Assert - Then
	s.Nil(err)

	var (
		id, accountFromId, accountToId string
		amount                         float64
		createdAt                      time.Time
	)

	stmt, err := s.db.Prepare("SELECT id, created_at, amount, account_from_id, account_to_id FROM transactions WHERE id = ?")
	s.Nil(err)
	defer stmt.Close()

	row := stmt.QueryRow(transaction.ID)
	err = row.Scan(&id, &createdAt, &amount, &accountFromId, &accountToId)

	s.Nil(err)
	s.Equal(transaction.ID, id)
	s.NotEmpty(createdAt)
	s.Equal(transaction.Amount, amount)
	s.Equal(transaction.AccountFrom.ID, accountFromId)
	s.Equal(transaction.AccountTo.ID, accountToId)

}
