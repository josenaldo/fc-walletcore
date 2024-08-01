package database

import (
	"database/sql"
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuite struct {
	suite.Suite
	ClientDb *ClientDb
	db       *sql.DB
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	db.Exec("CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, created_at DATETIME, updated_at DATETIME, name VARCHAR(255), email VARCHAR(255))")

	s.ClientDb = NewClientDb(db)

}

func (s *ClientDbTestSuite) TearDownSuite() {
	defer s.db.Close()

	s.db.Exec("DROP TABLE clients")
}

func TestClientDbTestSuit(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestSave() {
	client, _ := entity.NewClient("Zé Galinha", "ze@galinha.com")
	err := s.ClientDb.Save(client)
	s.Nil(err)
}

func (s *ClientDbTestSuite) TestGet() {
	client, _ := entity.NewClient("Zé Galinha", "ze@galinha.com")
	s.ClientDb.Save(client)

	clientFromDb, err := s.ClientDb.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientFromDb.ID)
	s.NotNil(clientFromDb.CreatedAt)
	s.NotNil(clientFromDb.UpdatedAt)
	s.Equal(client.Name, clientFromDb.Name)
	s.Equal(client.Email, clientFromDb.Email)
}
