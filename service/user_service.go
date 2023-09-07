//go:generate mockery --name=Service
package service

import (
	"gotest/db"
	"gotest/middleware"
	"gotest/types"

	"github.com/google/uuid"
)

type UserError struct{}

type ServiceStruct struct {
	dataStore db.DataStore
}

type Service interface {
	GetAllUsers() []types.User
	GetUserIdByUsername(username string) (userId string, err error)
	FindUserByUsernameAndPassword(username string, password string) (user types.User, err error)
	CreateNewUser(user types.User) types.User
	GetUserByName(name string) (user types.User, err error)
	UpdateUser(updatedUser types.User) (user types.User, err error)
	DeleteUserByName(username string) (user types.User, err error)
	GetAllGroceries() (groceries []types.Grocery)
	SearchGroceriesFromUser(query string) (groceries []types.Grocery)
	GetAllGroceriesFromUser(userId string) (groceries []types.Grocery, err error)
	CreateGroceryForUser(username string, groceryId int) (err error)
	CreateGrocery(grocery *types.Grocery) (groceryNew types.Grocery, err error)
	DeleteGroceryForUser(username string, groceryId int) (err error)
	GetGroceryByName(name string) (groceries []types.Grocery)
	UpdateGrocery(grocery types.Grocery) (updatedGrocery types.Grocery, errStatus error)
	UpdateStatusOfGrocery(groceryId int, status bool) (newGrocery types.Grocery, err error)
	DeleteGroceryById(id int) (grocery types.Grocery, errStatus error)
	GetRoleIdByName(username string) (roleId int, err error)
	GetUserByRole(role string) (users []types.User, err error)
}

func NewServiceStruct(dataStore db.DataStore) *ServiceStruct {
	return &ServiceStruct{
		dataStore: dataStore,
	}
}

func (m *UserError) Error() string {
	return "user not found"
}

func (us *ServiceStruct) GetAllUsers() []types.User {
	usersList := us.dataStore.GetAllUsers()
	return usersList
}

func (us *ServiceStruct) GetUserIdByUsername(username string) (userId string, err error) {
	var user types.User
	user, err = us.dataStore.GetUserByUsername(username)
	return user.ID.String(), err
}

func (us *ServiceStruct) CreateNewUser(user types.User) types.User {
	var createdUser types.User
	user.Password = string(middleware.HashString(user.Password))
	user.ID = uuid.New()
	user.RoleID = 2
	createdUser, _ = us.dataStore.CreateNewuser(user)
	return createdUser
}

func (us *ServiceStruct) GetUserByName(name string) (user types.User, err error) {
	userFound, err := us.dataStore.GetUserByUsername(name)
	if err != nil {
		return types.User{}, err
	} else if userFound.Name == "" {
		return types.User{}, &UserError{}
	}
	return userFound, nil
}

func (us *ServiceStruct) GetUserByRole(role string) (users []types.User, err error) {
	users, err = us.dataStore.GetUserByRole(role)
	if err != nil {
		return []types.User{}, err
	}
	return users, err
}

func (us *ServiceStruct) UpdateUser(updatedUser types.User) (user types.User, err error) {
	newUser, err := us.dataStore.UpdateUser(updatedUser)
	if err != nil {
		return types.User{}, err
	} else if newUser.Name == "" {
		return types.User{}, &UserError{}
	}
	return newUser, nil
}

func (us *ServiceStruct) FindUserByUsernameAndPassword(username string, password string) (user types.User, err error) {
	user, err = us.dataStore.FindUserByUsernameAndPassword(username, password)
	return
}

func (us *ServiceStruct) DeleteUserByName(username string) (user types.User, err error) {
	var deletedUser types.User
	deletedUser, err = us.dataStore.DeleteUserByName(username)
	if err != nil {
		return types.User{}, err
	} else if deletedUser.Age == 0 {
		return types.User{}, &UserError{}
	}
	return deletedUser, nil
}

func (us *ServiceStruct) GetRoleIdByName(username string) (roleId int, err error) {
	roleId, err = us.dataStore.GetRoleIdByName(username)
	return
}
