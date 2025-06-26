package interfaces

import (
	"GinEchoCrud/internal/models"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, int, error)
	GetUsersByID(ID int) (models.User, int, error)
	CreateUser(requestUser models.User) (models.User, int, error)
	UpdateUser(userID int, user models.UserRequests) (int, error)
	DeleteUser(userID int) (int, error)
}

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, int, error)
	GetUsersByID(ID int) (models.User, int, error)
	CreateUser(requestUser models.User) (models.User, int, error)
	UpdateUser(userID int, user models.UserRequests) (int, error)
	DeleteUser(userID int) (int, error)
}
