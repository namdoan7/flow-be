package service

import (
	"go-be/internal/common/utils"
	"go-be/internal/modules/user/model"
	"go-be/internal/modules/user/repository"
)

type Service interface {
	Register(name, email, password string) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(name, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return s.repo.Create(user)
}
