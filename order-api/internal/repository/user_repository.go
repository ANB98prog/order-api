package repository

import (
	goerr "errors"
	"github.com/ANB98prog/order-api/pkg/db"
	"github.com/ANB98prog/order-api/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(phone string) (*User, error)
	GetById(id uint) (*User, error)
	GetByPhone(phone string) (*User, error)
}

type userRepository struct {
	db *db.Db
}

func NewUserRepository(db *db.Db) UserRepository {
	return &userRepository{db: db}
}

var _ UserRepository = (*userRepository)(nil)

func (repo *userRepository) Create(phone string) (*User, error) {
	dbUser := User{Phone: phone}
	result := repo.db.Create(&dbUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dbUser, nil
}

func (repo *userRepository) GetById(id uint) (*User, error) {
	var user User
	result := repo.db.Take(&user, id)
	if result.Error != nil {
		if goerr.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repo *userRepository) GetByPhone(phone string) (*User, error) {
	var user User
	result := repo.db.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		if goerr.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
