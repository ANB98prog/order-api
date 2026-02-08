package service

import "github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/repository"

type UserService interface {
	Create(phone string) (*User, error)
	GetById(id uint) (*User, error)
	GetByPhone(phone string) (*User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

var _ UserService = (*userService)(nil)

func (s *userService) Create(phone string) (*User, error) {
	user, err := s.repo.Create(phone)
	if err != nil {
		return nil, err
	}

	return &User{Id: user.Id, Phone: phone}, nil
}

func (s *userService) GetById(id uint) (*User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return &User{Id: user.Id, Phone: user.Phone}, nil
}

func (s *userService) GetByPhone(phone string) (*User, error) {
	user, err := s.repo.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	return &User{Id: user.Id, Phone: user.Phone}, nil
}
