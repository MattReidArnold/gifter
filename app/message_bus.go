package app

import (
	"reflect"
)

type MessageType int

const (
	MessageTypeInvalid = iota
	MessageTypeEvent
	MessageTypeCommand
)

type Message interface {
	MessageType() MessageType
	Payload() interface{}
}

type HandlerFunc func(Message) error

type HandlerRegistry map[reflect.Type][]HandlerFunc
type messageBus struct {
	handlers HandlerRegistry
}
type message struct {
	messageType MessageType
	payload     interface{}
}

func NewCommandMessage(p interface{}) Message {
	return &message{
		messageType: MessageTypeCommand,
		payload:     p,
	}
}

func (m *message) MessageType() MessageType {
	return m.messageType
}

func (m *message) Payload() interface{} {
	return m.payload
}

type MessageBus interface {
	Handle(Message) error
	Register(reflect.Type, HandlerFunc)
}

func NewMessageBus() MessageBus {
	return &messageBus{
		handlers: HandlerRegistry{},
	}
}

func (mb *messageBus) Handle(m Message) error {
	handlers := mb.handlers[reflect.TypeOf(m.Payload())]
	for _, h := range handlers {
		err := h(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mb *messageBus) Register(t reflect.Type, h HandlerFunc) {
	handlers := mb.handlers[t]
	handlers = append(handlers, h)
	mb.handlers[t] = handlers
}
