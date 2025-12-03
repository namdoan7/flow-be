package event

import (
	"go-be/internal/event/dispatcher"
	_ "go-be/internal/event/handlers"

	"gorm.io/gorm"
)

type Dispatcher struct {
	*dispatcher.Dispatcher
	db *gorm.DB
}

func NewDispatcher(db *gorm.DB) *Dispatcher {
	return &Dispatcher{
		Dispatcher: dispatcher.GetDispatcher(),
		db:         db,
	}
}

// Emit event với data
func (d *Dispatcher) Emit(eventName string, data interface{}) {
	handlers := d.GetHandler(eventName)
	for _, h := range handlers {
		go h(data, d.db) // truyền db vào handler
	}
}
