package event

import "time"

var (
	BALANCE_UPDATED_NAME = "Balance.updated"
)

type BalanceUpdated struct {
	Name      string
	Payload   interface{}
	CreatedAt time.Time
}

func NewBalanceUpdated() *BalanceUpdated {
	return &BalanceUpdated{
		Name:      BALANCE_UPDATED_NAME,
		CreatedAt: time.Now(),
	}
}

func (e *BalanceUpdated) GetName() string {
	return e.Name
}

func (e *BalanceUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *BalanceUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *BalanceUpdated) GetDateTime() time.Time {
	return e.CreatedAt
}
