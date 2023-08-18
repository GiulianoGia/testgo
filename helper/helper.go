package helper

import (
	"gotest/db"
	"gotest/types"
)

func GetAllGroceries() (groceries []types.Grocery) {
	var groceryList []types.Grocery
	var err = db.DB.Find(&groceryList).Error
	if err != nil {
		return []types.Grocery{}
	}
	return groceryList
}

func CreateGrocery(grocery *types.Grocery) (groceryNew *types.Grocery, err error) {
	var newGrocery *types.Grocery
	err = db.DB.Create(&grocery).Find(&newGrocery).Error
	if err != nil {
		return newGrocery, err
	}
	return newGrocery, nil
}

func CreateGroceryForUser(username string, groceryId int) (err error) {
	var user types.User
	err = db.DB.Where("name = ?", username).First(&user).Error
	if err != nil {
		return
	}
	var userGrocery = types.UserGrocery{UserID: user.ID, GroceryID: groceryId}
	err = db.DB.Create(&userGrocery).Error
	if err != nil {
		return
	}
	return nil
}

func GetGroceryByName(name string) (groceries []types.Grocery) {
	var groceriesList = []types.Grocery{}
	var err = db.DB.Where("name = ?", name).Find(&groceriesList).Error
	if err != nil {
		return []types.Grocery{}
	}
	return groceriesList
}

func UpdateGrocery(grocery types.Grocery) (updatedGrocery types.Grocery, errStatus error) {
	var newGrocery = types.Grocery{}
	var err = db.DB.Save(&grocery).Where("id = ?", grocery.ID).Find(&newGrocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return newGrocery, nil
}

func DeleteGroceryById(id int) (grocery types.Grocery, errStatus error) {
	var oldGrocery = types.Grocery{}
	var err = db.DB.Where("id = ?", id).Find(&oldGrocery).Error
	if err != nil {
		return types.Grocery{}, err
	} else {
		err = db.DB.Delete(&oldGrocery).Error
		if err != nil {
			return types.Grocery{}, err
		}
		return oldGrocery, nil
	}
}
