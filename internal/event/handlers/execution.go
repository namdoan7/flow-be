package handlers

import (
	"fmt"
	"go-be/internal/event/dispatcher"

	"gorm.io/gorm"
)

type ExecutionData struct {
	ID   int
	Name string
}

func init() {
	dispatcher.GetDispatcher().RegisterHandler("flow.execution", handleExecutionCreated)
}

func handleExecutionCreated(data any, db *gorm.DB, event dispatcher.Emitter) {
	fmt.Printf("[Handler] ExecutionData")
	event.Emit("UserCreatedEvent", map[string]interface{}{
		"ID":   101,
		"Name": "Alice",
	})
}
