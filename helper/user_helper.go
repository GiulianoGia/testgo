package helper

import (
	"gotest/db"
	"gotest/types"
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
	err := db.DB.Create(user).Find(&createdUser).Error
	if err != nil {
		return types.User{}
	}
	return createdUser
}

func DeleteUserByName(userName string) (err error) {
	var deletedUser types.User
	err = db.DB.Where("name = ?", userName).Find(&deletedUser).Delete(&deletedUser).Error
	if err != nil {
		return err
	} else if deletedUser.Age == 0 {
		return &UserError{}
	}
	return nil
}
