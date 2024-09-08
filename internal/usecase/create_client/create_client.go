package create_client

import (
	"time"

	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/josenaldo/fc-walletcore/internal/gateway"
)

type CreateClientInputDto struct {
	Name  string
	Email string
}

type CreateClientOutputDto struct {
	ID        string
	CreatedAt time.Time
	UpdateAt  time.Time
	Name      string
	Email     string
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (usecase *CreateClientUseCase) Execute(input *CreateClientInputDto) (*CreateClientOutputDto, error) {

	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = usecase.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDto{
		ID:        client.ID,
		CreatedAt: client.CreatedAt,
		UpdateAt:  client.UpdatedAt,
		Name:      client.Name,
		Email:     client.Email,
	}, nil
}
