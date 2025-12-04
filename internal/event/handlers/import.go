package handlers

// Import để chạy init() trong các package con
import (
	_ "go-be/internal/event/handlers/execution"
	_ "go-be/internal/event/handlers/user"
)
