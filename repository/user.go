package repository

import (
	"inventory-manajemen-system/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(u entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(id int) (entity.User, error)
	Get() ([]entity.User, error)
	Update(user entity.User) (entity.User, error)
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

func (r *repository) FindByID(id int) (entity.User, error) {
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

func (r *repository) Update(user entity.User) (entity.User, error) {
	var updateUser = entity.InputUpdate{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}

	if result := r.db.Model(&user).Updates(updateUser); result != nil {
		return user, nil
	}

	return user, r.db.Error

}

func (r *repository) Delete(id int) (entity.User, error) {
	user := entity.User{}
	if result := r.db.Where("id = ?", id).Delete(&user); result != nil {
		return user, nil
	}
	return user, r.db.Error

}
