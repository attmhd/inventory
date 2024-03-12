package service

import (
	"errors"
	"inventory-manajemen-system/entity"
	"inventory-manajemen-system/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	AddUser(input entity.UserInput) (entity.User, error)
	GetUser() ([]entity.User, error)
	UpdateUser(update entity.InputUpdate) (entity.User, error)
	DeleteUser(id int)
	Login(input entity.LoginInput) (entity.User, error)
	IsEmailAvail(input entity.CheckEmailInput) (bool, error)
}

type service struct {
	repository repository.UserRepo
}

func NewService(repository repository.UserRepo) *service {
	return &service{repository}
}

func (s *service) GetUser() ([]entity.User, error) {
	result, err := s.repository.Get()
	if err != nil {
		return nil, err
	}

	var user []entity.User

	for _, value := range result {
		u := entity.User{
			Id:       value.Id,
			Username: value.Username,
			Password: value.Password,
			Email:    value.Email,
			Role:     value.Role,
		}
		user = append(user, u)
	}

	return user, nil
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

func (s *service) UpdateUser(update entity.InputUpdate) (entity.User, error) {
	user, err := s.repository.FindByID(update.Id)
	if err != nil {
		return user, err
	}

	user.Username = update.Username

	result, err := s.repository.Update(user)
	if err != nil {
		return result, err
	}

	return result, nil

}

func (s *service) DeleteUser(id int) {
	s.repository.Delete(id)
}

func (s *service) Login(input entity.LoginInput) (entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, err
	}

	return user, nil

}

func (s *service) IsEmailAvail(input entity.CheckEmailInput) (bool, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	if user.Id != 0 {
		return false, err
	}

	return true, nil

}
