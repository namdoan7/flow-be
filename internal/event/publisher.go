package event

import (
	"go-be/internal/common/utils"
	"go-be/internal/event/bus"
	_ "go-be/internal/event/handlers"
	"go-be/internal/event/types"
	"log"
	"reflect"
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

func (d *Publisher) GetEventItemDescription(v interface{}) map[string]string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	desc := map[string]string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		desc[f.Name] = f.Tag.Get("desc")
	}

	utils.PrintJSON("ItemDescription", desc)
	return desc
}

// Emit event với data
func (d *Publisher) Emit(eventName string, data interface{}) {
	handlers := d.dispatcher.GetHandler(eventName)
	for _, h := range handlers {
		go func(handler types.HandlerFunc) {
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
		go func(handler types.HandlerFunc) {
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

// EmitArray emits events sequentially (one by one)
func (d *Publisher) EmitArray(events []types.EventItem) {
	for _, evt := range events {
		d.EmitSync(evt.Name, evt.Data)
	}
}

// EmitArrayParallel emits all events in parallel
func (d *Publisher) EmitArrayParallel(events []types.EventItem) {
	var wg sync.WaitGroup

	for _, evt := range events {
		wg.Add(1)

		go func(e types.EventItem) {
			defer wg.Done()
			d.EmitSync(e.Name, e.Data)
		}(evt)
	}

	wg.Wait()
}

func (d *Publisher) Execution(flowId string, data any) {
	d.Emit("UserCreatedEvent", map[string]interface{}{
		"ID":   1012,
		"Name": "Alice",
	})
}
