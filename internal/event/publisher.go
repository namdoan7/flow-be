package event

import (
	"go-be/internal/event/bus"
	"go-be/internal/event/eventcore"
	_ "go-be/internal/event/handlers"
	"log"
	"runtime/debug"
	"sync"

	"gorm.io/gorm"
)

type Publisher struct {
	dispatcher *bus.Dispatcher
	db         *gorm.DB
}

func NewPublisher(db *gorm.DB) *Publisher {
	return &Publisher{
		dispatcher: bus.GetDispatcher(),
		db:         db,
	}
}

func (d *Publisher) GetDB() *gorm.DB {
	return d.db
}

// Emit event với data
func (d *Publisher) Emit(eventName string, data interface{}) {
	handlers := d.dispatcher.GetHandler(eventName)
	for _, h := range handlers {
		go func(handler eventcore.HandlerFunc) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Panic recovered in event handler: %v\n%s", r, debug.Stack())
				}
			}()
			handler(data, d) // truyền db vào handler
		}(h)
	}
}

// EmitSync emits an event and waits for all handlers to complete.
func (d *Publisher) EmitSync(eventName string, data interface{}) {
	handlers := d.dispatcher.GetHandler(eventName)
	var wg sync.WaitGroup
	for _, h := range handlers {
		wg.Add(1)
		go func(handler eventcore.HandlerFunc) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Panic recovered in event handler: %v\n%s", r, debug.Stack())
				}
			}()
			handler(data, d)
		}(h)
	}
	wg.Wait()
}
