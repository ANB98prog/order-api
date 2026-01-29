package repository

import "gorm.io/gorm"

type UserRepository interface {
	Create(user User) (User, error)
	GetById(id uint) (User, error)
	GetByPhone(id uint) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

var _ UserRepository = (*userRepository)(nil)

func (userRepository *userRepository) Create(user User) (User, error) {
	return User{}, nil
}

func (userRepository *userRepository) GetById(id uint) (User, error) {
	return User{}, nil
}

func (userRepository *userRepository) GetByPhone(id uint) (User, error) {
	return User{}, nil
}
