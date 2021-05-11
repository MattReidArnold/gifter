package app

import (
	"context"

	"github.com/mattreidarnold/gifter/domain"
)

type GroupRepository interface {
	Get(context.Context, string) (domain.Group, error)
	Add(context.Context, domain.Group) error
	Save(context.Context, domain.Group) error
}
