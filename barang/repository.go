package barang

import "gorm.io/gorm"

type Repository interface {
	Save(agd Admgd) (Admgd, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(agd Admgd) (Admgd, error) {
	err := r.db.Create(&agd).Error
	if err != nil {
		return agd, err
	}

	return agd, nil
}
