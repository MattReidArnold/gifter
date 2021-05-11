package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/interface/presenters"
)

func MakeAddGifter(msgBus app.MessageBus) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(presenters.AddGifterRequest)
		cmd := app.NewCommandMessage(domain.AddGifterCommand{
			GroupID: req.GroupID,
			Name:    req.Name,
		})
		err := msgBus.Handle(ctx, cmd)
		if err != nil {
			return nil, err
		}
		return presenters.AddGifterResponse(req), nil
	}
}
