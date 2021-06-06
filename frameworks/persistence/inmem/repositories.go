package inmem

import (
	"context"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
)

type GroupRepo struct {
	Groups map[string]domain.Group
}

func NewGroupRepository(groups ...domain.Group) app.GroupRepository {
	m := make(map[string]domain.Group, len(groups))
	for _, g := range groups {
		m[g.ID()] = g
	}
	return &GroupRepo{
		Groups: m,
	}
}

func (r *GroupRepo) Get(ctx context.Context, id string) (domain.Group, error) {
	g, ok := r.Groups[id]
	if !ok {
		return nil, app.ErrGroupNotFound
	}
	return g, nil
}

func (r *GroupRepo) Add(ctx context.Context, group domain.Group) error {
	if _, ok := r.Groups[group.ID()]; ok {
		return app.ErrGroupIDAlreadyExists
	}
	r.Groups[group.ID()] = group
	return nil
}

func (r *GroupRepo) Save(ctx context.Context, group domain.Group) error {
	if _, ok := r.Groups[group.ID()]; !ok {
		return app.ErrGroupNotFound
	}
	r.Groups[group.ID()] = group
	return nil
}
