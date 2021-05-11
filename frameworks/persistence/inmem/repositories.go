package persistence

import (
	"context"
	"errors"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
)

type groupRepo struct {
	groups map[string]domain.Group
	app.Logger
}

func NewGroupRepository(logger app.Logger, groups ...domain.Group) app.GroupRepository {
	m := make(map[string]domain.Group, len(groups))
	for _, g := range groups {
		m[g.ID()] = g
	}
	return &groupRepo{
		groups: m,
		Logger: logger,
	}
}

func (r *groupRepo) Get(ctx context.Context, id string) (domain.Group, error) {
	g, ok := r.groups[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return g, nil
}

func (r *groupRepo) Add(ctx context.Context, group domain.Group) error {
	if _, ok := r.groups[group.ID()]; ok {
		return errors.New("already exists")
	}

	r.groups[group.ID()] = group
	return nil
}

func (r *groupRepo) Save(ctx context.Context, group domain.Group) error {
	r.Logger.Info("Saving group", group)
	return nil
}
