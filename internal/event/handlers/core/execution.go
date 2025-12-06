package handlers

import (
	"go-be/internal/common/utils"
	"go-be/internal/event/bus"
	"go-be/internal/event/types"

	"github.com/mitchellh/mapstructure"
)

type IData struct {
	FlowId string      `json:"FlowId" desc:"Tên sự kiện (VD: UserCreatedEvent)"`
	Data   interface{} `json:"data" desc:"Tên sự kiện (VD: UserCreatedEvent)"`
}

var EventDesc string = "aaaaaa"
var EventName string = "flow.execution"

func init() {
	bus.GetDispatcher().RegisterHandler(bus.RegisterHandler{
		Doc: func() *types.DocumentFuncType {
			return &types.DocumentFuncType{Name: EventName, Desc: EventDesc, Fields: utils.TypeToJson(IData{})}
		},
		EventName: EventName,
		Handle:    HandleExecutionCreated,
	})
}

func HandleExecutionCreated(data any, emitter types.Emitter) {
	var eData IData
	mapstructure.Decode(data, &eData)
	emitter.Execution(eData.FlowId, eData.Data)
}
