package gateway

import (
	"errors"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

var (
	ErrorAccountNotFound   = errors.New("account not found")
	ErrorAccountExists     = errors.New("account already exists")
	ErrorAccountSaveFailed = errors.New("error saving account")
)

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
