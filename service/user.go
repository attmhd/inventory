package service

import (
	"inventory-manajemen-system/entity"
	"inventory-manajemen-system/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	AddUser(input entity.UserInput) (entity.User, error)
}

type service struct {
	repository repository.UserRepo
}

func NewService(repository repository.UserRepo) *service {
	return &service{repository}
}

func (s *service) AddUser(input entity.UserInput) (entity.User, error) {
	user := entity.User{}

	user.Id = input.Id
	user.Username = input.Username
	user.Email = input.Email
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(pwdHash)
	user.Role = input.Role

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
