package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, factory RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	Unregister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, factory RepositoryFactory) {
	u.Repositories[name] = factory
}

func (u *Uow) Unregister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {

	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	repositoryFactory := u.Repositories[name]
	if repositoryFactory == nil {
		return nil, fmt.Errorf("repository factory not found: %s", name)
	}

	repo := repositoryFactory(u.Tx)

	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	u.Tx = tx

	err = fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil

	return nil
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("commit error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	u.Tx = nil
	return nil
}
