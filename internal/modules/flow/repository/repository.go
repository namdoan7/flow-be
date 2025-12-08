package repository

import (
	"go-be/internal/modules/flow/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetPaginateFlow() ([]model.Flow, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetPaginateFlow() ([]model.Flow, error) {
	var flows []model.Flow
	err := r.db.Find(&flows).Error
	return flows, err
}
