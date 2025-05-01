package service

import "github.com/yourusername/simple-todo-golang/internal/models"

type UserService interface {
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id int) error
	GetUserByEmail(email string) (*models.User, error)
}
