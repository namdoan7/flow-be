package types

import "gorm.io/gorm"

type Emitter interface {
	Execution(executionId string, data interface{})
	Emit(eventName string, data interface{})
	EmitSync(eventName string, data interface{})
	GetDB() *gorm.DB
}

type DocumentFuncType struct {
	Name   string
	Desc   string
	Fields interface{}
}

type HandlerFunc func(data interface{}, emitter Emitter)
type DocumentFunc func() *DocumentFuncType

type EventItem struct {
	Name string
	Data interface{}
}
