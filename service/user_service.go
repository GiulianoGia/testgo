//go:generate mockery --name=UserServiceInterface
package service

import (
	"gotest/db"
	"gotest/middleware"
	"gotest/types"

	"github.com/google/uuid"
)

type UserError struct{}

type UserService struct {
	userRepo db.UserRepository
}

type UserServiceInterface interface {
	GetAllUsers() []types.User
	GetUserIdByUsername(username string) (userId string, err error)
	FindUserByUsernameAndPassword(username string, password string) (user types.User, err error)
	CreateNewUser(user types.User) types.User
	GetUserByName(name string) (user types.User, err error)
	UpdateUser(updatedUser types.User) (user types.User, err error)
	DeleteUserByName(username string) (user types.User, err error)
}

func NewUserService(ds *db.MariaDBDataStore, userRepo db.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (m *UserError) Error() string {
	return "user not found"
}

func (us *UserService) GetAllUsers() []types.User {
	usersList := us.userRepo.GetAllUsers()
	return usersList
}

func (us *UserService) GetUserIdByUsername(username string) (userId string, err error) {
	var user types.User
	user, err = us.userRepo.GetUserByUsername(username)
	return user.ID.String(), err
}

func (us *UserService) CreateNewUser(user types.User) types.User {
	var createdUser types.User
	user.Password = string(middleware.HashString(user.Password))
	user.ID = uuid.New()
	createdUser = us.userRepo.CreateNewuser(user)
	return createdUser
}

func (us *UserService) GetUserByName(name string) (user types.User, err error) {
	userFound, err := us.userRepo.GetUserByUsername(name)
	if err != nil {
		return types.User{}, err
	} else if userFound.Name == "" {
		return types.User{}, &UserError{}
	}
	return userFound, nil
}

func (us *UserService) UpdateUser(updatedUser types.User) (user types.User, err error) {
	newUser, err := us.userRepo.UpdateUser(updatedUser)
	if err != nil {
		return types.User{}, err
	} else if newUser.Name == "" {
		return types.User{}, &UserError{}
	}
	return newUser, nil
}

func (us *UserService) FindUserByUsernameAndPassword(username string, password string) (user types.User, err error) {
	user, err = us.userRepo.FindUserByUsernameAndPassword(username, password)
	return
}

func (us *UserService) DeleteUserByName(username string) (user types.User, err error) {
	var deletedUser types.User
	err = us.userRepo.DeleteUserByName(username)
	if err != nil {
		return types.User{}, err
	} else if deletedUser.Age == 0 {
		return types.User{}, &UserError{}
	}
	return deletedUser, nil
}
