package handlers

import (
	"go-be/internal/common/utils"
	"go-be/internal/event/bus"
	"go-be/internal/event/types"

	"github.com/mitchellh/mapstructure"
)

type IData struct {
	ID   int
	Name string
}

var EventDesc string = ""
var EventName string = "UserCreatedEvent"

func init() {
	bus.GetDispatcher().RegisterHandler(bus.RegisterHandler{
		Doc: func() *types.DocumentFuncType {
			return &types.DocumentFuncType{Name: EventName, Desc: EventDesc, Fields: utils.TypeToJson(IData{})}
		},
		EventName: EventName,
		Handle:    handleUserCreated,
	})
}

func handleUserCreated(data any, emitter types.Emitter) {
	var user IData
	mapstructure.Decode(data, &user)

	utils.PrintJSON("UserCreatedEvent", user)
}
