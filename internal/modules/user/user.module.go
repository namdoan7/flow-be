package user

import "gorm.io/gorm"

type Module struct {
	Handler    *Handler
	Service    Service
	Repository Repository
}

func New(db *gorm.DB) *Module {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{
		Handler:    handler,
		Service:    service,
		Repository: repo,
	}
}
