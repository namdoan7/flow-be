package repository

import (
	"go-be/internal/modules/user/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
