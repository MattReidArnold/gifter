package app

import "github.com/mattreidarnold/gifter/domain"

type GroupRepository interface {
	Get(string) (domain.Group, error)
	Add(domain.Group) error
}
