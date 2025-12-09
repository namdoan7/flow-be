package service

import (
	"go-be/internal/modules/flow/dto"
	"go-be/internal/modules/flow/model"
	"go-be/internal/modules/flow/repository"
)

type Service interface {
	GetPaginate() ([]model.Flow, error)
	UpdateFlow(id string, req dto.UpdateFlowRequest) (*model.Flow, error)
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

func (s *service) UpdateFlow(id string, req dto.UpdateFlowRequest) (*model.Flow, error) {
	flow, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if flow == nil {
		return nil, nil
	}
	flow.Name = req.Name
	flow.Status = req.Status

	if err := s.repo.Update(id, flow); err != nil {
		return nil, err
	}

	return flow, nil
}
