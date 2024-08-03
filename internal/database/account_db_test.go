package database

import (
	"database/sql"
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	AccountDb *AccountDB
	ClientDb  *ClientDb
	Client    *entity.Client
	db        *sql.DB
}

func (s *AccountDbTestSuite) SetupSuite() {
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

	s.ClientDb = NewClientDb(db)
	s.AccountDb = NewAccountDB(db)

	s.Client, _ = entity.NewClient("Zé Galinha", "ze@galinha.com")
	s.ClientDb.Save(s.Client)

}

func (s *AccountDbTestSuite) TearDownSuite() {
	defer s.db.Close()

	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDbTestSuit(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) TestSave() {
	// Arrange - Given
	account, _ := entity.NewAccount(s.Client)

	// Act - When
	err := s.AccountDb.Save(account)

	// Assert - Then
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestGet() {
	// Arrange - Given
	account, _ := entity.NewAccount(s.Client)
	s.AccountDb.Save(account)

	// Act - When
	accountFromDb, err := s.AccountDb.Get(account.ID)

	// Assert - Then
	s.Nil(err)
	s.Equal(account.ID, accountFromDb.ID)
	s.NotNil(accountFromDb.CreatedAt)
	s.NotNil(accountFromDb.UpdatedAt)
	s.Equal(account.Balance, accountFromDb.Balance)
	s.Equal(account.Client.ID, accountFromDb.Client.ID)
	s.Equal(account.Client.Name, accountFromDb.Client.Name)
	s.Equal(account.Client.Email, accountFromDb.Client.Email)

}
