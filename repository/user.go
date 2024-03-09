package repository

import "inventory-manajemen-system/model"

type UserRepo interface {
	Add(u model.User)
	Save(u model.User)
	Update(u model.User)
	Delete(u model.User)
}
