package createtransaction

import (
	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
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
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: transactionGateway,
		accountGateway:     accountGateway,
	}
}

func (usecase *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	accountFrom, err := usecase.accountGateway.FindByID(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := usecase.accountGateway.FindByID(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = usecase.transactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	err = usecase.accountGateway.Save(accountFrom)
	if err != nil {
		return nil, err
	}

	err = usecase.accountGateway.Save(accountTo)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDto{
		Id: transaction.ID,
	}, nil
}
