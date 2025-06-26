package service

import (
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"github.com/labstack/gommon/log"
)

type userService struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewUserService(userRepository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAllUsers() ([]models.User, int, error) {
	users, statusCode, err := s.userRepository.GetAllUsers()
	if err != nil {
		log.Fatal("error while retrieving users", err)
		return nil, statusCode, err
	}

	return users, statusCode, nil
}

func (s *userService) GetUsersByID(ID int) (models.User, int, error) {
	var (
		users      models.User
		err        error
		statusCode int
	)
	if ID == 0 {
		return users, helpers.StatusBadRequest, helpers.ErrMsgIdIsZero
	}

	users, statusCode, err = s.userRepository.GetUsersByID(ID)
	if err != nil {
		return users, statusCode, err
	}

	return users, statusCode, nil
}

func (s *userService) CreateUser(requestUser models.User) (models.User, int, error) {

	responseData, statusCode, responseError := s.userRepository.CreateUser(requestUser)
	if statusCode == 404 {
		return responseData, statusCode, helpers.ErrMsgNotFound
	}
	if responseError != nil {
		return responseData, statusCode, responseError
	}

	return responseData, statusCode, responseError
}

func (s *userService) UpdateUser(userID int, user models.UserRequests) (int, error) {
	statusCode, responseError := s.userRepository.UpdateUser(userID, user)
	if statusCode == 404 {
		return statusCode, helpers.ErrMsgNotFound
	}
	if responseError != nil {
		return statusCode, responseError
	}

	return statusCode, responseError
}

func (s *userService) DeleteUser(userID int) (int, error) {
	_, statusCodeResponse, errResponse := s.GetUsersByID(userID)
	if statusCodeResponse == 404 {
		return statusCodeResponse, helpers.ErrMsgNotFound
	}

	if errResponse != nil {
		return statusCodeResponse, errResponse
	}

	statusCode, responseError := s.userRepository.DeleteUser(userID)
	if statusCode == 404 {
		return statusCode, helpers.ErrMsgNotFound
	}
	if responseError != nil {
		return statusCode, responseError
	}

	return statusCode, responseError
}
