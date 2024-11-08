package event

import "time"

var (
	TRANSACTION_CREATED_NAME = "Transaction.created"
)

type TransactionCreated struct {
	Name      string
	Payload   interface{}
	CreatedAt time.Time
}

func NewTransactionCreated() *TransactionCreated {
	return &TransactionCreated{
		Name:      TRANSACTION_CREATED_NAME,
		CreatedAt: time.Now(),
	}
}

func (e *TransactionCreated) GetName() string {
	return e.Name
}

func (e *TransactionCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *TransactionCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *TransactionCreated) GetDateTime() time.Time {
	return e.CreatedAt
}
