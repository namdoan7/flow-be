package order

import "gorm.io/gorm"

type Repository interface {
	Create(order *Order) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(order *Order) error {
	return r.db.Create(order).Error
}
