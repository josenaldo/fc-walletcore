package create_account

import (
	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientId string `json:"client_id"`
}

type CreateAccountOutputDto struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (usecase *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := usecase.ClientGateway.Get(input.ClientId)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}

	err = usecase.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDto{
		ID: account.ID,
	}, nil
}
