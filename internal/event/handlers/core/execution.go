package handlers

import (
	"go-be/internal/event/bus"
	"go-be/internal/event/types"

	"github.com/mitchellh/mapstructure"
)

type ExecutionData struct {
	FlowId string      `json:"FlowId" desc:"Tên sự kiện (VD: UserCreatedEvent)"`
	Data   interface{} `json:"data" desc:"Tên sự kiện (VD: UserCreatedEvent)"`
}

func init() {
	bus.GetDispatcher().RegisterHandler("flow.execution", handleExecutionCreated)
}

func handleExecutionCreated(data any, emitter types.Emitter) {
	var eData ExecutionData
	mapstructure.Decode(data, &eData)

	emitter.Execution(eData.FlowId, eData.Data)
	emitter.GetEventItemDescription(ExecutionData{})
}
