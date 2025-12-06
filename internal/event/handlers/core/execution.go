package handlers

import (
	"go-be/internal/event/bus"
	"go-be/internal/event/types"
)

type ExecutionData struct {
	ID   int
	Name string
}

func init() {
	bus.GetDispatcher().RegisterHandler("flow.execution", handleExecutionCreated)
}

func handleExecutionCreated(data any, emitter types.Emitter) {
	emitter.Execution("abc", data)
}
