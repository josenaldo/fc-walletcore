package events

import "slices"

// EventDispatcher define um despachante de eventos.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Register registra um manipulador para um nome de evento específico.
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if ed.Has(eventName, handler) {
		return ErrorHandlerAlreadyRegistered
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// Remove remove um manipulador registrado para um nome de evento específico.
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {

	if !ed.Has(eventName, handler) {
		return nil
	}

	for i, h := range ed.handlers[eventName] {
		if h == handler {

			ed.handlers[eventName] = slices.Delete(ed.handlers[eventName], i, i+1)
			break
		}
	}

	return nil
}

// Has verifica se um manipulador está registrado para um nome de evento específico.
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {

	eventHandlers, exists := ed.handlers[eventName]

	if !exists {
		return false
	}

	for _, h := range eventHandlers {
		if h == handler {
			return true
		}
	}

	return false
}

// Clear remove todos os manipuladores registrados.
func (ed *EventDispatcher) Clear() {

	ed.handlers = make(map[string][]EventHandlerInterface)
}

// Dispatch despacha um evento para os manipuladores registrados.
func (ed *EventDispatcher) Dispatch(event EventInterface) error {

	handlers, exists := ed.handlers[event.GetName()]

	if !exists {
		return nil
	}

	for _, handler := range handlers {
		go handler.Handle(event)
	}

	return nil
}
