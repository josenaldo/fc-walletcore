package create_account

import (
	"testing"
	"time"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
	"github.com/josenaldo/fc-walletcore/internal/utils/assertions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUseCaseExecute(t *testing.T) {
	// Arrange - Given
	clientGatewayMock := &ClientGatewayMock{}

	clientGatewayMock.On("Get", mock.Anything).Return(&entity.Client{
		ID:        "321",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Zé Galinha",
		Email:     "ze@galinha.com",
	}, nil)
	clientGatewayMock.On("Save", mock.Anything).Return(nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	intput := &CreateAccountInputDto{
		ClientId: "123",
	}

	// Act - Whens
	output, err := uc.Execute(intput)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	assertions.IsUUID(t, output.ID)

	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreateAccountUseCaseExecuteWhenClientIsNotFound(t *testing.T) {
	// Arrange - Given
	clientGateway := &ClientGatewayMock{}

	clientGateway.On("Get", mock.Anything).Return(nil, gateway.ErrorClientNotFound)

	accountGateway := &AccountGatewayMock{}

	uc := NewCreateAccountUseCase(accountGateway, clientGateway)
	intput := &CreateAccountInputDto{
		ClientId: "123",
	}

	// Act - Whens
	output, err := uc.Execute(intput)

	// Assert - Then
	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, gateway.ErrorClientNotFound, err)
	clientGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Get", 1)
	clientGateway.AssertNumberOfCalls(t, "Save", 0)
	accountGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "Save", 0)
}
