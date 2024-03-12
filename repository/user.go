package repository

import (
	"inventory-manajemen-system/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(u entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(id string) (entity.User, error)
	Get() ([]entity.User, error)
	//Update(id int) (entity.User, error)
	Delete(id int) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get() ([]entity.User, error) {
	var u []entity.User
	if err := r.db.Find(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(id string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Save(u entity.User) (entity.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// func (r *repository) Update(i entity.UserInput) (entity.User, error) {
// 	user = new(entity.User)

// }

func (r *repository) Delete(id int) (entity.User, error) {
	user := entity.User{}
	if result := r.db.Where("id = ?", id).Delete(&user); result != nil {
		return user, nil
	} else {
		return user, r.db.Error
	}

}
