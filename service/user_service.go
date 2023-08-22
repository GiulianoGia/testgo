package service

import (
	"gotest/db"
	"gotest/middleware"
	"gotest/types"

	"github.com/google/uuid"
)

type UserError struct{}

func (m *UserError) Error() string {
	return "user not found"
}

func GetAllUsers() []types.User {
	var usersList []types.User
	var err = db.DB.Find(&usersList).Error
	if err != nil {
		return []types.User{}
	}
	return usersList
}

func CreateNewUser(user types.User) types.User {
	var createdUser types.User
	user.Password = string(middleware.HashString(user.Password))
	user.ID = uuid.New()
	err := db.DB.Create(user).Where("id", user.ID).Find(&createdUser).Error
	if err != nil {
		return types.User{}
	}
	return createdUser
}

func GetUserByName(name string) (user types.User, err error) {
	var userFound types.User
	err = db.DB.Where("name = ?", name).Find(&userFound).Error
	if err != nil {
		return types.User{}, err
	} else if userFound.Name == "" {
		return types.User{}, &UserError{}
	}
	return userFound, nil
}

func UpdateUser(updatedUser types.User) (user types.User, err error) {
	var newUser types.User
	err = db.DB.Save(&updatedUser).Where("id = ?", updatedUser.ID).Find(&newUser).Error
	if err != nil {
		return types.User{}, err
	} else if newUser.Name == "" {
		return types.User{}, &UserError{}
	}
	return newUser, nil
}

func DeleteUserByName(userName string) (user types.User, err error) {
	var deletedUser types.User
	err = db.DB.Where("name = ?", userName).Find(&deletedUser).Delete(&deletedUser).Error
	if err != nil {
		return types.User{}, err
	} else if deletedUser.Age == 0 {
		return types.User{}, &UserError{}
	}
	return deletedUser, nil
}
