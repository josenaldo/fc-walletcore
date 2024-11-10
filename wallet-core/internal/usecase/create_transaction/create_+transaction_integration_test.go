package create_transaction

import (
	"context"
	"database/sql"
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/database"
	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
	"github.com/josenaldo/fc-walletcore/internal/testutils"
	"github.com/josenaldo/fc-walletcore/pkg/events"
	"github.com/josenaldo/fc-walletcore/pkg/uow"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CreateTransactionUseCaseTestSuite struct {
	suite.Suite
	ctx                   context.Context
	testUtils             *testutils.TestDb
	clientrepository      gateway.ClientGateway
	accountrepository     gateway.AccountGateway
	transactionrepository gateway.TransactionGateway
	uow                   uow.UowInterface
	usecase               CreateTransactionUseCase
	eventDispatcher       *events.EventDispatcher

	client1 *entity.Client
	client2 *entity.Client

	toAccount   *entity.Account
	fromAccount *entity.Account
}

func (s *CreateTransactionUseCaseTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.testUtils = testutils.SetupTestDB(s.T())

	s.uow = uow.NewUow(s.ctx, s.testUtils.Db)

	s.uow.Register("ClientDB", func(tx *sql.Tx) interface{} {
		return database.NewClientDbWithTx(tx)
	})

	s.uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDbWithTx(tx)
	})

	s.uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDbWithTx(tx)
	})

	s.eventDispatcher = events.NewEventDispatcher()

	s.usecase = *NewCreateTransactionUseCase(s.uow, *s.eventDispatcher)

	s.clientrepository = database.NewClientDb(s.testUtils.Db)
	s.accountrepository = database.NewAccountDb(s.testUtils.Db)
	s.transactionrepository = database.NewTransactionDb(s.testUtils.Db)

	s.client1, _ = entity.NewClient("Ze Galinha", "ze@galinha.com")
	s.client2, _ = entity.NewClient("Maria Galinha", "maria@galinha.com")

	err := s.clientrepository.Save(s.client1)
	s.Require().NoError(err)
	err = s.clientrepository.Save(s.client2)
	s.Require().NoError(err)

	s.fromAccount, _ = entity.NewAccount(s.client2)
	s.fromAccount.Credit(300)
	s.toAccount, _ = entity.NewAccount(s.client1)
	s.toAccount.Credit(200)

	err = s.accountrepository.Save(s.fromAccount)
	s.Require().NoError(err)
	err = s.accountrepository.Save(s.toAccount)
	s.Require().NoError(err)

}

func (s *CreateTransactionUseCaseTestSuite) TearDownSuite() {
	s.testUtils.TearDownTestDB()
}

func TestCreateTransactionUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTransactionUseCaseTestSuite))
}

func (s *CreateTransactionUseCaseTestSuite) TestCreateTransactionUseCaseExecute() {
	// Arrange - Given
	input := CreateTransactionInputDto{
		FromAccountId: s.fromAccount.ID.String(),
		ToAccountId:   s.toAccount.ID.String(),
		Amount:        50,
	}

	// Act - When
	output, err := s.usecase.Execute(s.ctx, input)

	// Assert - Then
	s.NotNil(output)
	s.Nil(err)
	s.Equal(input.FromAccountId, output.FromAccountId)
	s.Equal(input.ToAccountId, output.ToAccountId)
	s.Equal(input.Amount, output.Amount)

	fromAccount, err := s.accountrepository.Get(s.fromAccount.ID)
	s.Require().NoError(err)

	toAccount, err := s.accountrepository.Get(s.toAccount.ID)
	s.Require().NoError(err)

	s.Equal(250.0, fromAccount.Balance)
	s.Equal(250.0, toAccount.Balance)
}

func (s *CreateTransactionUseCaseTestSuite) TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountFromNotFound() {
	// Arrange - Given
	input := CreateTransactionInputDto{
		FromAccountId: entity.NewEntityID().String(),
		ToAccountId:   s.toAccount.ID.String(),
		Amount:        50,
	}

	// Act - When
	output, err := s.usecase.Execute(s.ctx, input)

	// Assert - Then
	s.Nil(output)
	s.NotNil(err)
	s.EqualError(err, gateway.ErrorAccountNotFound.Error())

}

func (s *CreateTransactionUseCaseTestSuite) TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountToNotFound() {
	// Arrange - Given
	input := CreateTransactionInputDto{
		FromAccountId: s.fromAccount.ID.String(),
		ToAccountId:   entity.NewEntityID().String(),
		Amount:        50,
	}

	// Act - When
	output, err := s.usecase.Execute(s.ctx, input)

	// Assert - Then
	s.Nil(output)
	s.NotNil(err)
	s.EqualError(err, gateway.ErrorAccountNotFound.Error())

}

func (s *CreateTransactionUseCaseTestSuite) TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountFromHasInsufficientFunds() {
	// Arrange - Given

	input := CreateTransactionInputDto{
		FromAccountId: s.fromAccount.ID.String(),
		ToAccountId:   s.toAccount.ID.String(),
		Amount:        500,
	}

	// Act - When
	output, err := s.usecase.Execute(s.ctx, input)

	// Assert - Then
	s.Nil(output)
	s.NotNil(err)
	s.EqualError(err, entity.ErrorInsufficientFunds.Error())

}
