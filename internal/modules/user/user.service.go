package user

import (
	"go-be/internal/common/utils"
)

type Service interface {
	Register(name, email, password string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(name, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return s.repo.Create(user)
}
