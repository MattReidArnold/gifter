package handlers

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
)

func MakeAddGifter(d *app.Dependencies) (reflect.Type, app.HandlerFunc) {
	return reflect.TypeOf(domain.AddGifterCommand{}), func(ctx context.Context, msg app.Message) error {
		cmd, ok := msg.Payload().(domain.AddGifterCommand)
		if !ok {
			return errors.New("invalid message received by handler")
		}
		d.Logger.Info(fmt.Sprintf("Adding gifter %s to group %s", cmd.Name, cmd.GroupID))

		group, err := d.GroupRepository.Get(ctx, cmd.GroupID)
		if err != nil {
			return err
		}

		err = group.AddGifter(domain.NewGifter(cmd.Name))
		if err != nil {
			return err
		}
		d.Logger.Info((fmt.Sprintf("group: %+v", group)))
		err = d.GroupRepository.Save(ctx, group)
		return err
	}
}
