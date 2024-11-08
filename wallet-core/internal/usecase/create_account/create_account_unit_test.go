package create_account

import (
	"testing"
	"time"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
	"github.com/josenaldo/fc-walletcore/internal/testutils"
	"github.com/josenaldo/fc-walletcore/internal/testutils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCaseExecute(t *testing.T) {
	// Arrange - Given
	clientGatewayMock := &mocks.ClientGatewayMock{}

	clientGatewayMock.On("Get", mock.Anything).Return(&entity.Client{
		ID:        "321",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Zé Galinha",
		Email:     "ze@galinha.com",
	}, nil)

	accountGatewayMock := &mocks.AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	intput := CreateAccountInputDto{
		ClientId: "123",
	}

	// Act - Whens
	output, err := uc.Execute(intput)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	testutils.IsUUID(t, output.ID)

	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreateAccountUseCaseExecuteWhenClientIsNotFound(t *testing.T) {
	// Arrange - Given
	clientGateway := &mocks.ClientGatewayMock{}

	clientGateway.On("Get", mock.Anything).Return(nil, gateway.ErrorClientNotFound)

	accountGateway := &mocks.AccountGatewayMock{}

	uc := NewCreateAccountUseCase(accountGateway, clientGateway)
	intput := CreateAccountInputDto{
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
