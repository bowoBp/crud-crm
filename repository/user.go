package repository

import (
	"crud/entity"
	"fmt"
	"gorm.io/gorm"
	"math"
)

type User struct {
	db *gorm.DB
}

func NewUser(dbCrud *gorm.DB) User {
	return User{
		db: dbCrud,
	}

}

type UserInterfaceRepo interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserById(id uint) (entity.User, error)
	UpdateUser(user *entity.User, id uint) (any, error)
	DeleteUser(email string) (any, error)
	GetAllUser(page uint, username string) (uint, uint, int, uint, []entity.User, error)
}

// CreateUser new user
func (repo User) CreateUser(user *entity.User) (*entity.User, error) {
	err := repo.db.Model(&entity.User{}).Create(user).Error
	return user, err
}

// GetUserById get single user by id
func (repo User) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	repo.db.First(&user, "id = ? ", id)
	return user, nil
}

// UpdateUser multiple fields
func (repo User) UpdateUser(user *entity.User, id uint) (any, error) {
	fmt.Println(user.Name)
	err := repo.db.Debug().
		Model(&entity.User{}).
		Where("id = ?", id).
		Update("name", user.Name).Error

	//err := repo.db.Debug().Save(&user).Error
	return nil, err
}

// DeleteUser by Id and email
func (repo User) DeleteUser(email string) (any, error) {
	err := repo.db.Model(&entity.User{}).
		Where("email = ?", email).
		Delete(&entity.User{}).
		Error
	return nil, err
}

func (repo User) GetAllUser(page uint, username string) (uint, uint, int, uint, []entity.User, error) {
	var customers []entity.User
	var count int64
	var limit uint = 20
	var offset = limit * (page - 1)
	result := repo.db.Model(&entity.User{}).Count(&count)
	if result.Error != nil {
		// Handle the error
		return 0, 0, 0, 0, nil, result.Error
	}
	totalPages := uint(math.Ceil(float64(count) / float64(limit)))
	name := fmt.Sprintf("%%%s%%", username)

	err := repo.db.Select("*").
		Table("users").
		Select("*").
		Where("name LIKE ?", name).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&customers).
		Error
	if err != nil {
		return 0, 0, 0, 0, nil, err
	}
	return page, limit, int(count), totalPages, customers, nil
}
