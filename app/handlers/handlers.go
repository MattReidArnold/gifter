package handlers

import (
	"context"
	"reflect"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
)

func MakeAddGifter(d *app.Dependencies) (reflect.Type, app.HandlerFunc) {
	return reflect.TypeOf(domain.AddGifterCommand{}), func(ctx context.Context, msg app.Message) error {
		cmd, ok := msg.Payload().(domain.AddGifterCommand)
		if !ok {
			return app.ErrInvalidMessageTypeForHandler
		}

		return d.UseUnitOfWork(ctx, func(ctxUOW context.Context, uow app.UnitOfWork) error {
			group, err := uow.Groups().Get(ctxUOW, cmd.GroupID)
			if err != nil {
				return err
			}
			err = group.AddGifter(domain.NewGifter(cmd.GifterID, cmd.Name))
			if err != nil {
				return err
			}
			return uow.Groups().Save(ctxUOW, group)
		})
	}
}
