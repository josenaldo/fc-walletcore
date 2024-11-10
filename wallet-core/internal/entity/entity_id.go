package entity

import (
	"database/sql/driver"
	"errors"
	"strings"

	"github.com/oklog/ulid/v2"
)

var (
	ErrorEntityIDInvalid = errors.New("invalid id, must be a valid ULID")
	ErrorFailedToScan    = errors.New("failed to scan EntityID")
)

type EntityID struct {
	value ulid.ULID
}

func NewEntityID() EntityID {
	return EntityID{
		value: ulid.Make(),
	}
}

func ParseEntityID(s string) (EntityID, error) {
	id, err := ulid.Parse(strings.ToUpper(s))
	if err != nil {
		return EntityID{}, ErrorEntityIDInvalid
	}

	return EntityID{id}, nil
}

// String retorna a representação em string do EntityID
func (e EntityID) String() string {
	return e.value.String()
}

// Equals compara dois EntityIDs
func (e EntityID) Equals(other EntityID) bool {
	return e.value.Compare(other.value) == 0
}

// Scan implementa a interface sql.Scanner para leitura do banco de dados
func (e *EntityID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		return e.unmarshalString(v)
	case []byte:
		return e.unmarshalString(string(v))
	default:
		return ErrorFailedToScan
	}
}

func (e *EntityID) unmarshalString(s string) error {
	id, err := ulid.Parse(s)
	if err != nil {
		return ErrorEntityIDInvalid
	}
	e.value = id
	return nil
}

// Value implementa a interface driver.Valuer para gravação no banco de dados
func (e EntityID) Value() (driver.Value, error) {
	return e.String(), nil
}

// MarshalJSON implementa a interface json.Marshaler
func (e EntityID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON implementa a interface json.Unmarshaler
func (e *EntityID) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	id, err := ulid.Parse(str)
	if err != nil {
		return ErrorEntityIDInvalid
	}
	e.value = id
	return nil
}
