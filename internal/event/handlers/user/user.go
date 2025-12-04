package handlers

import (
	"go-be/internal/common/utils"
	"go-be/internal/event/bus"
	"go-be/internal/event/eventcore"

	"github.com/mitchellh/mapstructure"
)

type UserData struct {
	ID   int
	Name string
}

func init() {
	bus.GetDispatcher().RegisterHandler("UserCreatedEvent", handleUserCreated)
}

func handleUserCreated(data any, emitter eventcore.Emitter) {
	var user UserData
	mapstructure.Decode(data, &user)

	utils.PrintJSON("UserCreatedEvent", user)
}
