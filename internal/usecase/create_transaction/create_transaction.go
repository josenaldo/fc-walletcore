package createtransaction

import (
	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
	"github.com/josenaldo/fc-walletcore/pkg/events"
)

type CreateTransactionInputDto struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDto struct {
	Id string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
	EventDispatcher    events.EventDispatcher
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway, eventDispatcher events.EventDispatcher, transactionCreated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (usecase *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	accountFrom, err := usecase.AccountGateway.Get(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := usecase.AccountGateway.Get(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = usecase.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	err = usecase.AccountGateway.Save(accountFrom)
	if err != nil {
		return nil, err
	}

	err = usecase.AccountGateway.Save(accountTo)
	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDto{
		Id: transaction.ID,
	}

	usecase.TransactionCreated.SetPayload(output)
	usecase.EventDispatcher.Dispatch(usecase.TransactionCreated)

	return output, nil
}
