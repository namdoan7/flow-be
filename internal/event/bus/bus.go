package bus

import (
	"go-be/internal/event/types"
	"sync"
)

type Dispatcher struct {
	handlers  map[string][]types.HandlerFunc
	documents map[string]types.DocumentFunc
	mu        sync.RWMutex
}

type RegisterHandler struct {
	EventName string
	Doc       types.DocumentFunc
	Handle    types.HandlerFunc
}

var defaultDispatcher *Dispatcher

func init() {
	defaultDispatcher = &Dispatcher{
		handlers:  make(map[string][]types.HandlerFunc),
		documents: make(map[string]types.DocumentFunc),
	}
}

// Emit event với data
func (d *Dispatcher) GetHandler(eventName string) []types.HandlerFunc {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.handlers[eventName]
}

func (d *Dispatcher) GetDocument(eventName string) *types.DocumentFuncType {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.documents[eventName]()
}

// Register handler
func GetDispatcher() *Dispatcher {
	return defaultDispatcher
}

func (d *Dispatcher) RegisterHandler(data RegisterHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[data.EventName] = append(d.handlers[data.EventName], data.Handle)
	d.documents[data.EventName] = data.Doc
}
