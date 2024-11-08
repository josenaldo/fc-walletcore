package gateway

import (
	"errors"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

var (
	ErrorTransactionSaveFailed = errors.New("error saving transaction")
)

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
