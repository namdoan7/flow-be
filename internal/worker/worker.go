package worker

import (
	"log"

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
	log.Println("Worker started...")
	// Add worker logic here
}
