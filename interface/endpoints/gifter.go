package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/interface/presenters"
)

func MakeAddGifter(msgBus app.MessageBus) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(presenters.AddGifterRequest)
		cmd := app.NewCommandMessage(domain.AddGifterCommand{
			CircleID: req.CircleID,
			Name:     req.Name,
		})
		err := msgBus.Handle(cmd)
		if err != nil {
			return nil, err
		}
		return presenters.AddGifterResponse(req), nil
	}
}
