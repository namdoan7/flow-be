package auth

import "go-be/internal/modules/user"

type Module struct {
	Handler *Handler
	Service Service
}

func New(userRepo user.Repository) *Module {
	service := NewService(userRepo)
	handler := NewHandler(service)

	return &Module{
		Handler: handler,
		Service: service,
	}
}
