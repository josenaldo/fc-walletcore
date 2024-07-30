package createtransaction

import (
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
	"github.com/josenaldo/fc-walletcore/internal/utils/assertions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	client1 *entity.Client
	client2 *entity.Client

	accountTo   *entity.Account
	accountFrom *entity.Account

	accountGatewayMock     *AccountGatewayMock
	transactionGatewayMock *TransactionGatewayMock

	usecase *CreateTransactionUseCase
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func setupCreateTransactionUseCase() {
	client1, _ = entity.NewClient("Ze Galinha", "ze@galinha.com")
	client2, _ = entity.NewClient("Maria Galinha", "maria@galinha.com")

	accountTo, _ = entity.NewAccount(client1)
	accountTo.Credit(200)
	accountFrom, _ = entity.NewAccount(client2)
	accountFrom.Credit(300)

	accountGatewayMock = &AccountGatewayMock{}
	transactionGatewayMock = &TransactionGatewayMock{}

	usecase = NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)

}

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()
	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        50,
	}

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assertions.IsUUID(t, output.Id)
	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 2)
}

func TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountFromNotFound(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()
	input := CreateTransactionInputDto{
		AccountIdFrom: "non-existent-account",
		AccountIdTo:   accountTo.ID,
		Amount:        50,
	}
	accountGatewayMock.On("FindByID", "non-existent-account").Return(nil, gateway.ErrorAccountNotFound)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.EqualError(t, err, gateway.ErrorAccountNotFound.Error())
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 1)
	transactionGatewayMock.AssertNotCalled(t, "Create")
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 0)
}

func TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountToNotFound(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()

	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", "non-existent-account").Return(nil, gateway.ErrorAccountNotFound)
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   "non-existent-account",
		Amount:        50,
	}

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.EqualError(t, err, gateway.ErrorAccountNotFound.Error())
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNotCalled(t, "Create")
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 0)
}

func TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountFromHasInsufficientFunds(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()
	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        500,
	}

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.EqualError(t, err, entity.ErrorInsufficientFunds.Error())
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNotCalled(t, "Create")
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 0)
}

func TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountFromSaveFails(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()
	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)
	accountGatewayMock.On("Save", accountFrom).Return(gateway.ErrorAccountSaveFailed)

	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        50,
	}

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.EqualError(t, err, gateway.ErrorAccountSaveFailed.Error())
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNotCalled(t, "Create")
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreateTransactionUseCaseExecuteReturnErrorWhenAccountToSaveFails(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()
	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)
	accountGatewayMock.On("Save", accountFrom).Return(nil)
	accountGatewayMock.On("Save", accountTo).Return(gateway.ErrorAccountSaveFailed)

	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        50,
	}

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.EqualError(t, err, gateway.ErrorAccountSaveFailed.Error())
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNotCalled(t, "Create")
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 2)
}

func TestCreateTransactionUseCaseExecuteReturnErrorWhenTransactionCreateFails(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()
	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)
	accountGatewayMock.On("Save", accountFrom).Return(nil)
	accountGatewayMock.On("Save", accountTo).Return(nil)

	transactionGatewayMock.On("Create", mock.Anything).Return(gateway.ErrorTransactionSaveFailed)

	input := CreateTransactionInputDto{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        50,
	}

	// Act - When
	output, err := usecase.Execute(input)

	// Assert - Then
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.EqualError(t, err, gateway.ErrorTransactionSaveFailed.Error())
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 0)
}
