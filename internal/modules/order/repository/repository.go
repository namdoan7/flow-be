package repository

import (
	"go-be/internal/modules/order/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(order *model.Order) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}
