package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/interface/presenters"
)

func MakeAddGifter(d *app.Dependencies) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(presenters.AddGifterRequest)

		gifterID, err := d.GenerateID()
		if err != nil {
			return nil, err
		}

		cmd := app.NewCommandMessage(domain.AddGifterCommand{
			Name:     req.Name,
			GifterID: gifterID,
			GroupID:  req.GroupID,
		})

		err = d.MessageBus.Handle(ctx, cmd)
		if err != nil {
			return nil, err
		}

		return presenters.AddGifterResponse{
			GifterID: gifterID,
			GroupID:  req.GroupID,
			Name:     req.Name,
		}, nil
	}
}
