package tests

import (
	"GinEchoCrud/internal/models"
	"errors"
)

type MockUserService struct{}
type MockUserServiceFail struct{}

func (m *MockUserService) CreateUser(user models.User) (models.User, int, error) {
	user.ID = 99
	return user, 200, nil
}

func (m *MockUserServiceFail) CreateUser(user models.User) (models.User, int, error) {
	return models.User{}, 400, errors.New("error")
}

func (m *MockUserService) GetAllUsers() ([]models.User, int, error) {
	return []models.User{
		{ID: 14, Name: "UserTest1", Age: 44, Email: "userTest@email1.com"},
		{ID: 10, Name: "UserTest2", Age: 31, Email: "userTest@email2.com"},
		{ID: 19, Name: "UserTest3", Age: 19, Email: "userTest@email3.com"},
	}, 200, nil
}

func (m *MockUserServiceFail) GetAllUsers() ([]models.User, int, error) {
	return nil, 400, errors.New("error")
}

func (m *MockUserService) GetUsersByID(ID int) (models.User, int, error) {
	return models.User{
		ID:    14,
		Name:  "UserTest1",
		Age:   44,
		Email: "userTest@email1.com",
	}, 200, nil
}

func (m *MockUserServiceFail) GetUsersByID(ID int) (models.User, int, error) {
	return models.User{}, 400, nil
}
