package app

import (
	"context"

	"github.com/mattreidarnold/gifter/domain"
)

type Dependencies struct {
	GroupRepository
	GenerateID
	Logger
	MessageBus
}

type Logger interface {
	Info(args ...interface{})
	Error(err error, args ...interface{})
}

type GenerateID func() (string, error)

type GroupRepository interface {
	Get(context.Context, string) (domain.Group, error)
	Add(context.Context, domain.Group) error
	Save(context.Context, domain.Group) error
}
