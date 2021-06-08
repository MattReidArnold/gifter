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
	UseUnitOfWork
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

type UnitOfWork interface {
	Groups() GroupRepository
}

type UseUnitOfWork func(context.Context, func(context.Context, UnitOfWork) error) error
