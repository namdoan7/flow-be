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

func NewDispatcherInit() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[string][]HandlerFunc),
	}
}

var (
	defaultDispatcher *Dispatcher
	once              sync.Once
)

func NewDispatcher() *Dispatcher {
	once.Do(func() {
		defaultDispatcher = NewDispatcherInit()
	})
	return defaultDispatcher
}

// Emit event với data
func (d *Dispatcher) GetHandler(eventName string) []HandlerFunc {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.handlers[eventName]
}

// Register handler
func (d *Dispatcher) Register(eventName string, h HandlerFunc) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventName] = append(d.handlers[eventName], h)
}
