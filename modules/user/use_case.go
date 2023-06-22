package user

import (
	"crud/entity"
	"crud/repository"
	"time"
)

type UseCaseUser interface {
	CreateUser(user UserParam) (entity.User, error)
	GetUserById(id uint) (entity.User, error)
	UpdateUser(user UserParam, id uint) (any, error)
	DeleteUser(email string) (any, error)
	GetAllUser(page uint, username string) (uint, uint, int, uint, []entity.User, error)
}

type useCaseUser struct {
	userRepo repository.UserInterfaceRepo
}

func (uc useCaseUser) CreateUser(user UserParam) (entity.User, error) {
	var newUser *entity.User

	newUser = &entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		Phone:     user.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.userRepo.CreateUser(newUser)
	if err != nil {
		return *newUser, err
	}
	return *newUser, nil
}

func (uc useCaseUser) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	user, err := uc.userRepo.GetUserById(id)
	return user, err
}

func (uc useCaseUser) UpdateUser(user UserParam, id uint) (any, error) {
	var editUser *entity.User
	editUser = &entity.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		Phone:     user.Phone,
		UpdatedAt: time.Now(),
	}

	_, err := uc.userRepo.UpdateUser(editUser, id)
	if err != nil {
		return editUser, err
	}
	return editUser, nil
}

func (uc useCaseUser) DeleteUser(email string) (any, error) {
	_, err := uc.userRepo.DeleteUser(email)
	return nil, err
}

func (uc useCaseUser) GetAllUser(page uint, username string) (uint, uint, int, uint, []entity.User, error) {
	var customer []entity.User
	page, perPage, total, totalPages, customer, err := uc.userRepo.GetAllUser(page, username)
	return page, perPage, total, totalPages, customer, err
}
