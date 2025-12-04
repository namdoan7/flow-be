package bus

import (
	"sync"

	"go-be/internal/event/eventcore"
)

type Dispatcher struct {
	handlers map[string][]eventcore.HandlerFunc
	mu       sync.RWMutex
}

var defaultDispatcher *Dispatcher

func init() {
	defaultDispatcher = &Dispatcher{
		handlers: make(map[string][]eventcore.HandlerFunc),
	}
}

// Emit event với data
func (d *Dispatcher) GetHandler(eventName string) []eventcore.HandlerFunc {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.handlers[eventName]
}

// Register handler
func GetDispatcher() *Dispatcher {
	return defaultDispatcher
}

func (d *Dispatcher) RegisterHandler(eventName string, h eventcore.HandlerFunc) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventName] = append(d.handlers[eventName], h)
}
