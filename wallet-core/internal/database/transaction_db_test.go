package database

import (
	"testing"
	"time"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/testutils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	TransactionDb *TransactionDB

	clientFrom  *entity.Client
	accountFrom *entity.Account

	clientTo  *entity.Client
	accountTo *entity.Account

	testUtils *testutils.TestDb
}

func (s *TransactionDbTestSuite) SetupSuite() {
	s.testUtils = testutils.SetupTestDB(s.T())

	s.TransactionDb = NewTransactionDb(s.testUtils.Db)

	s.clientFrom, _ = entity.NewClient("Zé Galinha", "ze@galinha.com")
	s.testUtils.Db.Exec("INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)",
		s.clientFrom.ID, s.clientFrom.CreatedAt, s.clientFrom.UpdatedAt, s.clientFrom.Name, s.clientFrom.Email)

	s.clientTo, _ = entity.NewClient("Maria Galinha", "maria@galinha.com")
	s.testUtils.Db.Exec("INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)",
		s.clientTo.ID, s.clientTo.CreatedAt, s.clientTo.UpdatedAt, s.clientTo.Name, s.clientTo.Email)

	s.accountFrom, _ = entity.NewAccount(s.clientFrom)
	s.accountFrom.Credit(1000)
	s.testUtils.Db.Exec("INSERT INTO accounts (id, created_at, updated_at, balance, client_id) VALUES (?, ?, ?, ?, ?)",
		s.accountFrom.ID, s.accountFrom.CreatedAt, s.accountFrom.UpdatedAt, s.accountFrom.Balance, s.accountFrom.Client.ID)

	s.accountTo, _ = entity.NewAccount(s.clientTo)
	s.accountTo.Credit(1000)
	s.testUtils.Db.Exec("INSERT INTO accounts (id, created_at, updated_at, balance, client_id) VALUES (?, ?, ?, ?, ?)",
		s.accountTo.ID, s.accountTo.CreatedAt, s.accountTo.UpdatedAt, s.accountTo.Balance, s.accountTo.Client.ID)

}

func (s *TransactionDbTestSuite) TearDownSuite() {
	s.testUtils.TearDownTestDB()
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

	stmt, err := s.testUtils.Db.Prepare("SELECT id, created_at, amount, account_from_id, account_to_id FROM transactions WHERE id = ?")
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
