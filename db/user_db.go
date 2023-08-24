//go:generate mockery --name=UserRepository
package db

import "gotest/types"

type UserRepository interface {
	GetUserByUsername(username string) (user types.User, err error)
	GetAllUsers() (userList []types.User)
	CreateNewuser(user types.User) (createdUser types.User)
	UpdateUser(updatedUser types.User) (newUser types.User, err error)
	DeleteUserByName(username string) (err error)
	FindUserByUsernameAndPassword(username string, password string) (user types.User, err error)
}

func (ds *MariaDBDataStore) GetUserByUsername(username string) (user types.User, err error) {
	err = ds.db.Where("name = ?", username).First(&user).Error
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (ds *MariaDBDataStore) GetAllUsers() (userList []types.User) {
	var usersList []types.User
	var err = ds.db.Find(&usersList).Error
	if err != nil {
		return []types.User{}
	}
	return userList
}

func (ds *MariaDBDataStore) FindUserByUsernameAndPassword(username string, password string) (user types.User, err error) {
	err = ds.db.Where(&types.User{Name: username, Password: password}).Find(&user).Error
	if err != nil {
		return types.User{}, err
	}
	return
}

func (ds *MariaDBDataStore) CreateNewuser(user types.User) (createdUser types.User, err error) {
	err = ds.db.Create(user).Where("id", user.ID).Find(&createdUser).Error
	if err != nil {
		return types.User{}, err
	}
	return createdUser, nil
}

func (ds *MariaDBDataStore) UpdateUser(updatedUser types.User) (newUser types.User, err error) {
	err = ds.db.Save(&updatedUser).Where("id = ?", updatedUser.ID).Find(&newUser).Error
	if err != nil {
		return types.User{}, err
	}
	return newUser, nil
}

func (ds *MariaDBDataStore) DeleteUserByName(username string) (err error) {
	var deletedUser types.User
	err = ds.db.Where("name = ?", username).Find(&deletedUser).Delete(&deletedUser).Error
	return err
}
