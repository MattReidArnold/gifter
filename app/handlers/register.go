package handlers

import "github.com/mattreidarnold/gifter/app"

func RegisterAll(d *app.Dependencies) {
	//Commands
	d.MessageBus.RegisterCommandHandler(MakeAddGifter(d))
	//Events
	d.MessageBus.RegisterEventHandler(SendGifterWelcomeNotification(d))
}
