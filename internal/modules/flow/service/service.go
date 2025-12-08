package service

import (
	"go-be/internal/modules/flow/model"
	"go-be/internal/modules/flow/repository"
)

type Service interface {
	GetPaginate() ([]model.Flow, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetPaginate() ([]model.Flow, error) {
	return s.repo.GetPaginateFlow()
}
