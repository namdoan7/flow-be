package repository

import (
	"errors"
	"go-be/internal/modules/flow/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetPaginateFlow() ([]model.Flow, error)
	GetById(id string) (*model.Flow, error)
	Update(id string, flow *model.Flow) error
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

func (r *repository) GetById(id string) (*model.Flow, error) {
	var flow model.Flow
	err := r.db.Where("id = ?", id).First(&flow).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &flow, err
}

func (r *repository) Update(id string, flow *model.Flow) error {
	return r.db.Model(&model.Flow{}).
		Where("id = ?", id).
		Updates(flow).Error
}
