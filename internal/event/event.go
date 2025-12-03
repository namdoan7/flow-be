package event

import (
	"go-be/internal/event/dispatcher"
	_ "go-be/internal/event/handlers"
	"log"
	"runtime/debug"
	"sync"

	"gorm.io/gorm"
)

type Event struct {
	*dispatcher.Dispatcher
	db *gorm.DB
}

func NewEvent(db *gorm.DB) *Event {
	return &Event{
		Dispatcher: dispatcher.GetDispatcher(),
		db:         db,
	}
}

// Emit event với data
func (d *Event) Emit(eventName string, data interface{}) {
	handlers := d.GetHandler(eventName)
	for _, h := range handlers {
		go func(handler dispatcher.HandlerFunc) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Panic recovered in event handler: %v\n%s", r, debug.Stack())
				}
			}()
			handler(data, d.db) // truyền db vào handler
		}(h)
	}
}

// EmitSync emits an event and waits for all handlers to complete.
func (d *Event) EmitSync(eventName string, data interface{}) {
	handlers := d.GetHandler(eventName)
	var wg sync.WaitGroup
	for _, h := range handlers {
		wg.Add(1)
		go func(handler dispatcher.HandlerFunc) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Panic recovered in event handler: %v\n%s", r, debug.Stack())
				}
			}()
			handler(data, d.db)
		}(h)
	}
	wg.Wait()
}
