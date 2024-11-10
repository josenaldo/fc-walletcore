package gateway

import (
	"errors"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

var (
	ErrorClientNotFound = errors.New("client not found")
	ErrorClientExists   = errors.New("client already exists")
	ErrorClientSave     = errors.New("error saving client")
)

type ClientGateway interface {
	GetAll() ([]*entity.Client, error)
	Get(id entity.EntityID) (*entity.Client, error)
	Save(client *entity.Client) error
}
