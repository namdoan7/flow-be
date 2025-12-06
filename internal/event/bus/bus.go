package bus

import (
	"go-be/internal/event/types"
	"sync"
)

type Dispatcher struct {
	handlers map[string][]types.HandlerFunc
	mu       sync.RWMutex
}

var defaultDispatcher *Dispatcher

func init() {
	defaultDispatcher = &Dispatcher{
		handlers: make(map[string][]types.HandlerFunc),
	}
}

// Emit event với data
func (d *Dispatcher) GetHandler(eventName string) []types.HandlerFunc {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.handlers[eventName]
}

// Register handler
func GetDispatcher() *Dispatcher {
	return defaultDispatcher
}

func (d *Dispatcher) RegisterHandler(eventName string, h types.HandlerFunc) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventName] = append(d.handlers[eventName], h)
}
