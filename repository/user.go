package repository

import (
	"inventory-manajemen-system/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(u entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(u entity.User) (entity.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
