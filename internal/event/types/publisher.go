package types

import "gorm.io/gorm"

type Emitter interface {
	Execution(executionId string, data interface{})
	Emit(eventName string, data interface{})
	EmitSync(eventName string, data interface{})
	GetDB() *gorm.DB
	GetEventItemDescription(v any) map[string]string
}

type HandlerFunc func(data interface{}, emitter Emitter)

type EventItem struct {
	Name string
	Data interface{}
}
