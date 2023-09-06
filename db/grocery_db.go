//go:generate mockery --name=DataStore
package db

import (
	"gotest/types"

	"github.com/google/uuid"
)

type DataStore interface {
	GetAllGroceries() (groceries []types.Grocery, err error)
	GetAllGroceriesFromUser(userId string) (userGroceries []types.UserGrocery, err error)
	FindGroceryWithId(groceryId int) (grocery types.Grocery, err error)
	CreateGrocery(newGrocery types.Grocery) (grocery types.Grocery, err error)
	CreateGroceryForUser(userId uuid.UUID, groceryId int) (err error)
	DeleteGroceryFromUser(userId uuid.UUID, groceryId int) (err error)
	GetGroceriesByName(name string) (groceries []types.Grocery, err error)
	UpdateGrocery(grocery types.Grocery) (newGrocery types.Grocery, err error)
	UpdateStatusOfGrocery(groceryId int, status bool) (updatedGrocery types.Grocery, err error)
	DeleteGrocery(groceryId int) (grocery types.Grocery, err error)
	GetUserByUsername(username string) (user types.User, err error)
	GetAllUsers() (userList []types.User)
	CreateNewuser(user types.User) (createdUser types.User, er error)
	UpdateUser(updatedUser types.User) (newUser types.User, err error)
	DeleteUserByName(username string) (deletedUser types.User, err error)
	FindUserByUsernameAndPassword(username string, password string) (user types.User, err error)
	GetRoleIdByName(username string) (roleId int, err error)
	GetUserByRole(role string) (users []types.User, err error)
}

func (ds *MariaDBDataStore) GetAllGroceries() (groceries []types.Grocery, err error) {
	err = ds.db.Find(&groceries).Error
	if err != nil {
		return []types.Grocery{}, err
	}
	return
}

func (ds *MariaDBDataStore) GetGroceriesByName(name string) (groceries []types.Grocery, err error) {
	var groceriesList = []types.Grocery{}
	err = ds.db.Where("name = ?", name).Find(&groceriesList).Error
	if err != nil {
		return []types.Grocery{}, err
	}
	return groceriesList, nil
}

func (ds *MariaDBDataStore) GetAllGroceriesFromUser(userId string) (userGroceries []types.UserGrocery, err error) {
	err = ds.db.Where("user_id = ?", userId).Find(&userGroceries).Error
	if err != nil {
		return []types.UserGrocery{}, err
	}
	return
}

func (ds *MariaDBDataStore) FindGroceryWithId(groceryId int) (grocery types.Grocery, err error) {
	err = ds.db.Where("id = ?", groceryId).Find(&grocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return
}

func (ds *MariaDBDataStore) UpdateGrocery(grocery types.Grocery) (newGrocery types.Grocery, err error) {
	err = ds.db.Save(&grocery).Where("id = ?", grocery.ID).Find(&newGrocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return newGrocery, nil
}

func (ds *MariaDBDataStore) UpdateStatusOfGrocery(groceryId int, status bool) (updatedGrocery types.Grocery, err error) {
	err = ds.db.Model(&types.Grocery{}).Where("id = ?", groceryId).Update("done", status).Find(&updatedGrocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return updatedGrocery, nil
}

func (ds *MariaDBDataStore) CreateGrocery(newGrocery types.Grocery) (grocery types.Grocery, err error) {
	err = ds.db.Create(&newGrocery).Where("id = ?", newGrocery.ID).Find(&grocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return grocery, nil
}

func (ds *MariaDBDataStore) CreateGroceryForUser(userId uuid.UUID, groceryId int) (err error) {
	var userGrocery = types.UserGrocery{UserID: userId, GroceryID: groceryId}
	err = ds.db.Create(&userGrocery).Error
	return
}

func (ds *MariaDBDataStore) DeleteGroceryFromUser(userId uuid.UUID, groceryId int) (err error) {
	grocery := types.UserGrocery{UserID: userId, GroceryID: groceryId}
	err = ds.db.Delete(&grocery).Error
	return
}

func (ds *MariaDBDataStore) DeleteGrocery(groceryId int) (grocery types.Grocery, err error) {
	var oldGrocery types.Grocery
	err = ds.db.Where("id = ?", groceryId).Delete(&oldGrocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return oldGrocery, nil
}
