package database

import (
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/testutils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuite struct {
	suite.Suite
	ClientDb  *ClientDb
	testUtils *testutils.TestDb
}

func (s *ClientDbTestSuite) SetupSuite() {
	s.testUtils = testutils.SetupTestDB(s.T())

	s.ClientDb = NewClientDb(s.testUtils.Db)

}

func (s *ClientDbTestSuite) TearDownSuite() {
	s.testUtils.TearDownTestDB()
}

func TestClientDbTestSuit(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestSave() {
	// Arrange - Given
	client, _ := entity.NewClient("Zé Galinha", "ze@galinha.com")

	// Act - When
	err := s.ClientDb.Save(client)

	// Assert - Then
	s.Nil(err)
}

func (s *ClientDbTestSuite) TestGet() {
	// Arrange - Given
	client, _ := entity.NewClient("Zé Galinha", "ze@galinha.com")
	s.ClientDb.Save(client)

	// Act - When
	clientFromDb, err := s.ClientDb.Get(client.ID)

	// Assert - Then
	s.Nil(err)
	s.Equal(client.ID, clientFromDb.ID)
	s.NotNil(clientFromDb.CreatedAt)
	s.NotNil(clientFromDb.UpdatedAt)
	s.Equal(client.Name, clientFromDb.Name)
	s.Equal(client.Email, clientFromDb.Email)
}
