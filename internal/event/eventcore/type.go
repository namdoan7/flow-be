package eventcore

import "gorm.io/gorm"

type Emitter interface {
	Emit(eventName string, data interface{})
	EmitSync(eventName string, data interface{})
	GetDB() *gorm.DB
}

type HandlerFunc func(data interface{}, emitter Emitter)
