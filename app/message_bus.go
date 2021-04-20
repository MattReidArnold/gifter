package app

import (
	"container/list"
	"errors"
	"fmt"
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
	PayloadType() reflect.Type
}

type HandlerFunc func(Message) error

type EventHandlerRegistry map[reflect.Type][]HandlerFunc

type CommandHandlerRegistry map[reflect.Type]HandlerFunc
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

func NewEventMessage(p interface{}) Message {
	return &message{
		messageType: MessageTypeEvent,
		payload:     p,
	}
}

func (m *message) MessageType() MessageType {
	return m.messageType
}

func (m *message) Payload() interface{} {
	return m.payload
}

func (m *message) PayloadType() reflect.Type {
	return reflect.TypeOf(m.Payload())
}

func (m *message) String() string {
	return fmt.Sprintf("%v: %v", m.PayloadType(), m.Payload())
}

type MessageBus interface {
	Handle(Message) error
	RegisterCommandHandler(reflect.Type, HandlerFunc)
	RegisterEventHandler(reflect.Type, HandlerFunc)
}
type messageBus struct {
	commandHandlers CommandHandlerRegistry
	eventHandlers   EventHandlerRegistry
	queue           *list.List
	Logger
}

func NewMessageBus(l Logger) MessageBus {
	return &messageBus{
		commandHandlers: CommandHandlerRegistry{},
		eventHandlers:   EventHandlerRegistry{},
		queue:           list.New(),
		Logger:          l,
	}
}

func (mb *messageBus) Handle(m Message) error {
	mb.queue.Init().PushFront(m)
	for mb.queue.Len() > 0 {
		node := mb.queue.Front()
		msg, ok := node.Value.(Message)
		if !ok {
			return errors.New("invalid thing in message bus queue")
		}
		mb.Logger.Info(fmt.Sprintf("handling message %v", msg))
		var err error
		switch msg.MessageType() {
		case MessageTypeCommand:
			err = mb.handleCommand(msg)
		case MessageTypeEvent:
			err = mb.handleEvent(msg)
		default:
			err = errors.New("invalid message type")
		}
		if err != nil {
			return err
		}
		mb.queue.Remove(node)
	}
	return nil
}

func (mb *messageBus) handleCommand(m Message) error {
	handler, ok := mb.commandHandlers[m.PayloadType()]
	if !ok {
		return errors.New("command handler not found")
	}
	err := handler(m)
	if err != nil {
		return err
	}
	// Collect events from UOW
	return nil
}

func (mb *messageBus) handleEvent(m Message) error {
	handlers := mb.eventHandlers[m.PayloadType()]
	for _, h := range handlers {
		err := h(m)
		if err != nil {
			return err
		}
	}
	// Collect events from UOW
	return nil
}

func (mb *messageBus) RegisterEventHandler(t reflect.Type, h HandlerFunc) {
	handlers := mb.eventHandlers[t]
	handlers = append(handlers, h)
	mb.eventHandlers[t] = handlers
}

func (mb *messageBus) RegisterCommandHandler(t reflect.Type, h HandlerFunc) {
	if _, alreadyRegistered := mb.commandHandlers[t]; alreadyRegistered {
		panic(errors.New("multiple handlers registered for a single command type"))
	}
	mb.commandHandlers[t] = h
}
