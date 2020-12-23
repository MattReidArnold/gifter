package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/mattreidarnold/gifter/app/usecase"
	"github.com/mattreidarnold/gifter/domain/entities"
	"github.com/mattreidarnold/gifter/interface/presenters"
)

func MakeAddGifter(addUser usecase.AddGifter) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(presenters.AddGifterRequest)
		g, err := addUser.Execute(entities.Gifter{Name: req.Name})
		if err != nil {
			return nil, nil
		}
		return presenters.AddGifterResponse{
			Name: g.Name,
		}, nil
	}
}
