package create_transaction

import (
	"context"
	"testing"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/testutils/mocks"
	"github.com/josenaldo/fc-walletcore/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	client1 *entity.Client
	client2 *entity.Client

	accountTo   *entity.Account
	accountFrom *entity.Account

	uowMock    *mocks.UowMock
	ctx        context.Context
	usecase    *CreateTransactionUseCase
	dispatcher *events.EventDispatcher
)

func setupCreateTransactionUseCase() {
	client1, _ = entity.NewClient("Ze Galinha", "ze@galinha.com")
	client2, _ = entity.NewClient("Maria Galinha", "maria@galinha.com")

	accountFrom, _ = entity.NewAccount(client2)
	accountFrom.Credit(300)
	accountTo, _ = entity.NewAccount(client1)
	accountTo.Credit(200)

	uowMock = &mocks.UowMock{}
	uowMock.On("Do", mock.Anything, mock.Anything).Return(nil)
	dispatcher = events.NewEventDispatcher()
	ctx = context.Background()

	usecase = NewCreateTransactionUseCase(uowMock, *dispatcher)

}

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	// Arrange - Given
	setupCreateTransactionUseCase()

	input := CreateTransactionInputDto{
		FromAccountId: accountFrom.ID,
		ToAccountId:   accountTo.ID,
		Amount:        50,
	}

	// Act - When
	output, err := usecase.Execute(ctx, input)

	// Assert - Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	uowMock.AssertNumberOfCalls(t, "Do", 1)
}
