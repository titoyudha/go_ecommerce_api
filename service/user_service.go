package service

import (
	"database/sql"
	"log"

	"github.com/titoyudha/go_ecommerce_api/dto"
	"github.com/titoyudha/go_ecommerce_api/model"
	"github.com/titoyudha/go_ecommerce_api/repositories"
)

type UserService interface {
	GetUser() []model.User
	GetUserByID(userID string) model.User
	CreateNewUser(dto dto.UserRequestBody) (model.User, error)
	DeleteUser(userID string) error
}

type userService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserService {
	return &userService{
		UserRepository: userRepository,
	}
}

// CreateNewUser implements UserService
func (repo *userService) CreateNewUser(data dto.UserRequestBody) (model.User, error) {
	user := model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}

	err := repo.UserRepository.SignUp(user)
	if err != err {
		log.Panic(err)
		return model.User{}, sql.ErrTxDone
	}
	return user, nil
}

// DeleteUser implements UserService
func (repo *userService) DeleteUser(userID string) error {
	user := repo.GetUserByID(userID)
	if user.UserID == "" {
		return sql.ErrNoRows
	}

	return repo.DeleteUser(user.UserID)
}

// GetUser implements UserService
func (repo *userService) GetUser() []model.User {
	return repo.UserRepository.FindAll()
}

// GetUserByID implements UserService
func (repo *userService) GetUserByID(userID string) model.User {
	return repo.GetUserByID(userID)
}
