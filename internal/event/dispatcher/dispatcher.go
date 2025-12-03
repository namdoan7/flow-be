package dispatcher

import (
	"sync"

	"gorm.io/gorm"
)

type HandlerFunc func(data interface{}, db *gorm.DB)

type Dispatcher struct {
	handlers map[string][]HandlerFunc
	mu       sync.RWMutex
}

var defaultDispatcher *Dispatcher

func init() {
	defaultDispatcher = &Dispatcher{
		handlers: make(map[string][]HandlerFunc),
	}
}

// Emit event với data
func (d *Dispatcher) GetHandler(eventName string) []HandlerFunc {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.handlers[eventName]
}

// Register handler
func GetDispatcher() *Dispatcher {
	return defaultDispatcher
}

func (d *Dispatcher) RegisterHandler(eventName string, h HandlerFunc) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventName] = append(d.handlers[eventName], h)
}
