package mocks

import (
	"context"
	"sync"

	"github.com/josenaldo/fc-walletcore/pkg/uow"
	"github.com/stretchr/testify/mock"
)

type UowMock struct {
	mock.Mock
	wg *sync.WaitGroup
}

func (m *UowMock) Register(name string, factory uow.RepositoryFactory) {
	m.Called(name, factory)
}

func (m *UowMock) Unregister(name string) {
	m.Called(name)
}

func (m *UowMock) GetRepository(ctx context.Context, name string) (interface{}, error) {
	args := m.Called(ctx, name)
	return args.Get(0), args.Error(1)
}

func (m *UowMock) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
	args := m.Called(ctx, fn)
	if m.wg != nil {
		defer m.wg.Done()
	}

	return args.Error(0)
}

func (m *UowMock) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UowMock) CommitOrRollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UowMock) SetWaitGroup(wg *sync.WaitGroup) {
	m.wg = wg
}
