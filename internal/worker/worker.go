package worker

import (
	"log"

	"go-be/internal/event"

	"gorm.io/gorm"
)

type Worker struct {
	db *gorm.DB
}

func NewWorker(db *gorm.DB) *Worker {
	return &Worker{
		db: db,
	}
}

func (w *Worker) Run() {
	// Dispatcher singleton với DB injected
	dispatcher := event.NewPublisher(w.db)

	// Emit event
	dispatcher.Emit("flow.execution", map[string]interface{}{
		"ID":   101,
		"Name": "Alice",
	})

	log.Println("Worker started...")
	// Add worker logic here
}

func (w *Worker) GetExecution(executionId string) {

}
