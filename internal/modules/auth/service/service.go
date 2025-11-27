package service

import (
	"errors"
	"go-be/internal/common/utils"
	"go-be/internal/modules/user"
)

type Service interface {
	Login(email, password string) (string, error)
}

type service struct {
	userRepo user.Repository
}

func NewService(userRepo user.Repository) Service {
	return &service{userRepo: userRepo}
}

func (s *service) Login(email, password string) (string, error) {
	u, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, u.Password) {
		return "", errors.New("invalid credentials")
	}

	// TODO: Generate real JWT
	return "mock-jwt-token", nil
}
