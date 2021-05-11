package persistence

import (
	"errors"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
)

type groupRepo struct {
	groups map[string]domain.Group
}

func NewGroupRepository(groups ...domain.Group) app.GroupRepository {
	m := make(map[string]domain.Group, len(groups))
	for _, g := range groups {
		m[g.ID()] = g
	}
	return &groupRepo{
		groups: m,
	}
}

func (r *groupRepo) Get(id string) (domain.Group, error) {
	g, ok := r.groups[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return g, nil
}

func (r *groupRepo) Add(group domain.Group) error {
	if _, ok := r.groups[group.ID()]; ok {
		return errors.New("already exists")
	}

	r.groups[group.ID()] = group
	return nil
}
