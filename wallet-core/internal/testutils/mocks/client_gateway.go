package mocks

import (
	"github.com/josenaldo/fc-walletcore/internal/entity"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) GetAll() ([]*entity.Client, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Get(id entity.EntityID) (*entity.Client, error) {
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
