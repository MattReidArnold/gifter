package app

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
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

type HandlerFunc func(context.Context, Message) error

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
	return fmt.Sprintf("%v: %+v", m.PayloadType(), m.Payload())
}

type MessageBus interface {
	Handle(context.Context, Message) error
	EnqueueEvents(...interface{}) error
	RegisterCommandHandler(reflect.Type, HandlerFunc)
	RegisterEventHandler(reflect.Type, HandlerFunc)
}
type messageBus struct {
	commandHandlers CommandHandlerRegistry
	eventHandlers   EventHandlerRegistry
	queue           *list.List
	Logger
	sync.Mutex
}

func NewMessageBus(l Logger) MessageBus {
	return &messageBus{
		commandHandlers: CommandHandlerRegistry{},
		eventHandlers:   EventHandlerRegistry{},
		queue:           list.New(),
		Logger:          l,
	}
}

func (mb *messageBus) Handle(ctx context.Context, m Message) error {
	mb.initMessageQueue(m)
	for mb.hasMessages() {
		msg, err := mb.nextMessage()
		if err != nil {
			return err
		}
		mb.Logger.Info(fmt.Sprintf("handling message %v", msg))

		switch msg.MessageType() {
		case MessageTypeCommand:
			err = mb.handleCommand(ctx, msg)
		case MessageTypeEvent:
			err = mb.handleEvent(ctx, msg)
		default:
			err = errors.New("invalid message type")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (mb *messageBus) EnqueueEvents(events ...interface{}) error {
	mb.Lock()
	defer mb.Unlock()

	for _, e := range events {
		m := NewEventMessage(e)
		mb.queue.PushBack(m)
	}
	return nil
}

func (mb *messageBus) handleCommand(ctx context.Context, m Message) error {
	handler, ok := mb.commandHandlers[m.PayloadType()]
	if !ok {
		return errors.New("command handler not found")
	}
	err := handler(ctx, m)
	if err != nil {
		return err
	}
	// Collect events from UOW
	return nil
}

func (mb *messageBus) handleEvent(ctx context.Context, m Message) error {
	handlers := mb.eventHandlers[m.PayloadType()]
	for _, h := range handlers {
		err := h(ctx, m)
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

func (mb *messageBus) hasMessages() bool {
	mb.Lock()
	defer mb.Unlock()
	return mb.queue.Len() > 0
}

func (mb *messageBus) initMessageQueue(msg Message) {
	mb.Lock()
	defer mb.Unlock()
	mb.queue.Init().PushFront(msg)
}

func (mb *messageBus) nextMessage() (Message, error) {
	mb.Lock()
	defer mb.Unlock()

	node := mb.queue.Front()
	mb.queue.Remove(node)
	msg, ok := node.Value.(Message)
	if !ok {
		return nil, errors.New("invalid thing in message bus queue")
	}
	return msg, nil
}
