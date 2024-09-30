package database

import (
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/testutils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	AccountDb *AccountDB
	client    *entity.Client
	testUtils *testutils.TestDb
}

func (s *AccountDbTestSuite) SetupSuite() {
	s.testUtils = testutils.SetupTestDB(s.T())

	s.AccountDb = NewAccountDb(s.testUtils.Db)

	s.client, _ = entity.NewClient("Zé Galinha", "ze@galinha.com")
	s.testUtils.Db.Exec("INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)",
		s.client.ID, s.client.CreatedAt, s.client.UpdatedAt, s.client.Name, s.client.Email)

}

func (s *AccountDbTestSuite) TearDownSuite() {
	s.testUtils.TearDownTestDB()
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) TestSave() {
	// Arrange - Given
	account, _ := entity.NewAccount(s.client)

	// Act - When
	err := s.AccountDb.Save(account)

	// Assert - Then
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestGet() {
	// Arrange - Given
	account, _ := entity.NewAccount(s.client)
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

func (s *AccountDbTestSuite) TestGetNotFound() {
	// Arrange - Given
	account, _ := entity.NewAccount(s.client)
	s.AccountDb.Save(account)

	// Act - When
	accountFromDb, err := s.AccountDb.Get("non-existing-id")

	// Assert - Then
	s.Nil(accountFromDb)
	s.NotNil(err)
}

func (s *AccountDbTestSuite) TestUpdate() {
	// Arrange - Given
	account, _ := entity.NewAccount(s.client)
	s.AccountDb.Save(account)

	// Act - When
	account.Credit(100)
	err := s.AccountDb.UpdateBalance(account)

	// Assert - Then
	s.Nil(err)

	accountFromDb, _ := s.AccountDb.Get(account.ID)
	s.Equal(100.0, accountFromDb.Balance)
}
