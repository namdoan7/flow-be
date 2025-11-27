package order

import "github.com/google/uuid"

type Service interface {
	CreateOrder(userID uuid.UUID, total float64) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateOrder(userID uuid.UUID, total float64) error {
	order := &Order{
		UserID: userID,
		Total:  total,
		Status: "pending",
	}
	return s.repo.Create(order)
}
