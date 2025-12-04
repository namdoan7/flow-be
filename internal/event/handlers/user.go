package handlers

import (
	"fmt"
	"go-be/internal/event/dispatcher"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type UserData struct {
	ID   int
	Name string
}

func init() {
	dispatcher.GetDispatcher().RegisterHandler("UserCreatedEvent", handleUserCreated)
}

func handleUserCreated(data any, db *gorm.DB, emitter dispatcher.Emitter) {
	var user UserData
	mapstructure.Decode(data, &user) // decode bất kỳ map/object sang struct

	// if db != nil {
	// 	db.Create(&user)
	// }

	fmt.Printf("[Handler] User created: ID=%d, Name=%s\n", user.ID, user.Name)
}
