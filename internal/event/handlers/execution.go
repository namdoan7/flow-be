package handlers

import (
	"fmt"
	"go-be/internal/event/bus"
	"go-be/internal/event/eventcore"
)

type ExecutionData struct {
	ID   int
	Name string
}

func init() {
	bus.GetDispatcher().RegisterHandler("flow.execution", handleExecutionCreated)
}

func handleExecutionCreated(data any, emitter eventcore.Emitter) {
	fmt.Printf("[Handler] ExecutionData")
	emitter.Emit("UserCreatedEvent", map[string]interface{}{
		"ID":   101,
		"Name": "Alice",
	})
}
