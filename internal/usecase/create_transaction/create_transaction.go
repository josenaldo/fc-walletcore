package create_transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/event"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
	"github.com/josenaldo/fc-walletcore/pkg/events"
	"github.com/josenaldo/fc-walletcore/pkg/uow"
)

type CreateTransactionInputDto struct {
	FromAccountId string  `json:"from_account_id"`
	ToAccountId   string  `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	Id            string  `json:"id"`
	FromAccountId string  `json:"from_account_id"`
	ToAccountId   string  `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow             uow.UowInterface
	EventDispatcher events.EventDispatcher
}

func NewCreateTransactionUseCase(
	Uow uow.UowInterface,
	eventDispatcher events.EventDispatcher,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:             Uow,
		EventDispatcher: eventDispatcher,
	}
}

func (usecase *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	output := &CreateTransactionOutputDto{}

	err := usecase.Uow.Do(ctx, func(uow *uow.Uow) error {
		accountRepository := usecase.getAccountRepository(ctx)
		transactionRepository := usecase.getTransactionRepository(ctx)

		repo, err := usecase.Uow.GetRepository(ctx, "ClientDB")
		if err != nil {
			return err
		}
		clientRepository := repo.(gateway.ClientGateway)
		clients, err := clientRepository.GetAll()
		if err != nil {
			return err
		}

		if len(clients) == 0 {
			return errors.New("no clients found")
		}

		fmt.Println("Clients found: ", clients)

		accountFrom, err := accountRepository.Get(input.FromAccountId)
		if err != nil {
			return err
		}

		accountTo, err := accountRepository.Get(input.ToAccountId)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		output.Id = transaction.ID
		output.FromAccountId = transaction.AccountFrom.ID
		output.ToAccountId = transaction.AccountTo.ID
		output.Amount = transaction.Amount

		return nil

	})
	if err != nil {
		return nil, err
	}

	event := event.NewTransactionCreated()
	event.SetPayload(output)

	usecase.EventDispatcher.Dispatch(event)
	return output, nil
}

func (usecase *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := usecase.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (usecase *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := usecase.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
