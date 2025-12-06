package handlers

// Import để chạy init() trong các package con
import (
	_ "go-be/internal/event/handlers/core"
	_ "go-be/internal/event/handlers/user"
)
