package handlers

import "github.com/mattreidarnold/gifter/app"

func RegisterAll(d *app.Dependencies) {
	d.MessageBus.RegisterCommandHandler(MakeAddGifter(d))
}
