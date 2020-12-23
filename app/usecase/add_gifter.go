package usecase

import (
	"fmt"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain/entities"
)

type addGifter struct {
	logger app.Logger
}

type AddGifter interface {
	Execute(entities.Gifter) (entities.Gifter, error)
}

func NewAddGifter(d *app.Dependencies) AddGifter {
	return &addGifter{
		logger: d.Logger,
	}
}

func (a *addGifter) Execute(gifter entities.Gifter) (entities.Gifter, error) {
	a.logger.Info(fmt.Sprintf("Adding gifter %s", gifter.Name))
	return entities.Gifter{}, nil
}
