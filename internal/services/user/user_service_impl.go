package service

import (
	"github.com/yourusername/simple-todo-golang/internal/models"
	repository "github.com/yourusername/simple-todo-golang/internal/repository/user"
)

type UserServiceImpl struct {
	repo repository.UserRespositoryImpl
}

func NewUserService(repo repository.UserRespositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUserByID(id int) (*models.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *models.User) (*models.User, error) {
	user, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUser(user *models.User) (*models.User, error) {
	user, err := s.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
