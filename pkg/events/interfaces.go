package events

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrorHandlerAlreadyRegistered = errors.New("handler already registered for event name")
)

// EventInterface define a interface para um evento.
type EventInterface interface {

	// GetName retorna o nome do evento.
	GetName() string

	// GetDateTime retorna a data e hora do evento.
	GetDateTime() time.Time

	// GetPayload retorna a carga útil do evento.
	GetPayload() interface{}

	// SetPayload define a carga útil do evento.
	SetPayload(payload interface{})
}

// EventHandlerInterface define a interface para um manipulador de eventos.
type EventHandlerInterface interface {

	// Handle processa um evento.
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// EventDispatcherInterface define a interface para um despachante de eventos.
type EventDispatcherInterface interface {

	// Register registra um manipulador para um nome de evento específico.
	Register(eventName string, handler EventHandlerInterface) error

	// Remove remove um manipulador registrado para um nome de evento específico.
	Remove(eventName string, handler EventHandlerInterface) error

	// Has verifica se um manipulador está registrado para um nome de evento específico.
	Has(eventName string, handler EventHandlerInterface) bool

	// Clear remove todos os manipuladores registrados.
	Clear() error

	// Dispatch despacha um evento para os manipuladores registrados.
	Dispatch(event EventInterface) error
}
