package service

import (
	"go-be/internal/modules/order/model"
	"go-be/internal/modules/order/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateOrder(userID uuid.UUID, total float64) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateOrder(userID uuid.UUID, total float64) error {
	order := &model.Order{
		UserID: userID,
		Total:  total,
		Status: "pending",
	}
	return s.repo.Create(order)
}
