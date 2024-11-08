package create_client

import (
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/testutils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	defaultName  string
	defaultEmail string
)

func init() {
	defaultName = "Zé Galinha"
	defaultEmail = "ze@galinha.com"
}

func TestCreateClientUseCaseExecute(t *testing.T) {
	// Arrange - Given
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)
	intput := CreateClientInputDto{
		Name:  defaultName,
		Email: defaultEmail,
	}

	// Act - When
	output, err := uc.Execute(intput)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, intput.Name, output.Name)
	assert.Equal(t, intput.Email, output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}

func TestReturnErrorWhenCreateClientWithEmptyName(t *testing.T) {
	// Arrange - Given
	m := &mocks.ClientGatewayMock{}
	uc := NewCreateClientUseCase(m)
	intput := CreateClientInputDto{
		Name:  "",
		Email: defaultEmail,
	}

	// Act - When
	output, err := uc.Execute(intput)

	// Assert - Then
	assert.EqualError(t, err, entity.ErrorClientNameIsRequired.Error())
	assert.Nil(t, output)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 0)
}

func TestReturnErrorWhenCreateClientWithEmptyEmail(t *testing.T) {
	// Arrange - Given
	m := &mocks.ClientGatewayMock{}
	uc := NewCreateClientUseCase(m)
	intput := CreateClientInputDto{
		Name:  defaultName,
		Email: "",
	}

	// Act - When
	output, err := uc.Execute(intput)

	// Assert - Then
	assert.EqualError(t, err, entity.ErrorClientEmailIsRequired.Error())
	assert.Nil(t, output)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 0)
}
