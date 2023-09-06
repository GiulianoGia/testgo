package db

import (
	"gotest/types"
)

func (ds *MariaDBDataStore) GetUserByUsername(username string) (user types.User, err error) {
	err = ds.db.Preload("Role").Where("name = ?", username).First(&user).Error
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (ds *MariaDBDataStore) GetUserByRole(role string) (users []types.User, err error) {
	roleId := types.GetRoleId(role)
	err = ds.db.Preload("Role").Find(&users, "role_id = ?", roleId).Error
	if err != nil {
		return []types.User{}, err
	}
	return users, nil
}

func (ds *MariaDBDataStore) GetAllUsers() (userList []types.User) {
	var err = ds.db.Preload("Role").Find(&userList).Error
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

func (ds *MariaDBDataStore) DeleteUserByName(username string) (deletedUser types.User, err error) {
	err = ds.db.Where("name = ?", username).Find(&deletedUser).Delete(&deletedUser).Error
	return
}

func (ds *MariaDBDataStore) GetRoleIdByName(username string) (roleId int, err error) {
	var user types.User
	err = ds.db.Preload("Role").Where("name = ?", username).Find(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Role.ID, nil
}
