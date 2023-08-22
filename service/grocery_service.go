package service

import (
	"fmt"
	"gotest/db"
	"gotest/types"
	"strings"
)

func GetAllGroceries() (groceries []types.Grocery) {
	var groceryList []types.Grocery
	var err = db.DB.Find(&groceryList).Error
	if err != nil {
		return []types.Grocery{}
	}
	return groceryList
}

func SearchGroceriesFromUser(query string) (groceriesList []types.Grocery) {
	var groceries []types.Grocery
	groceries = append(groceries, types.Grocery{ID: 1, Name: "Test", Quantity: 12, Done: false})
	var groceriesLikeQuery []types.Grocery
	for _, grocery := range groceries {
		if strings.Contains(grocery.Name, query) {
			groceriesLikeQuery = append(groceriesLikeQuery, grocery)
		}
	}
	return groceriesLikeQuery
}

func GetAllGroceriesFromUser(userId string) (groceries []types.Grocery, err error) {
	var userGroceries []types.UserGrocery
	err = db.DB.Where("user_id = ?", userId).Find(&userGroceries).Error
	if err != nil {
		return []types.Grocery{}, err
	}
	for _, userGrocery := range userGroceries {
		var grocery types.Grocery
		err = db.DB.Where("id = ?", userGrocery.GroceryID).Find(&grocery).Error
		if err != nil {
			return []types.Grocery{}, err
		}
		groceries = append(groceries, grocery)
	}
	return groceries, nil
}

func GetUserIdByUsername(username string) (userId string, err error) {
	var user *types.User
	err = db.DB.Where("name = ?", username).Find(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
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

func DeleteGroceryForUser(username string, groceryId int) (err error) {
	deletedGrocery, err := DeleteGroceryById(groceryId)
	if err != nil {
		return
	}
	user, err := GetUserByName(username)
	if err != nil {
		return
	}
	grocery := types.UserGrocery{UserID: user.ID, GroceryID: int(deletedGrocery.ID)}
	err = db.DB.Delete(&grocery).Error
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

func UpdateStatusOfGrocery(groceryId int, status bool) (newGrocery types.Grocery, err error) {
	var grocery types.Grocery
	fmt.Println(groceryId, status)
	err = db.DB.Model(&types.Grocery{}).Where("id = ?", groceryId).Update("done", status).Find(&grocery).Error
	if err != nil {
		return types.Grocery{}, err
	}
	return grocery, nil
}

func DeleteGroceryById(id int) (grocery types.Grocery, errStatus error) {
	var oldGrocery = types.Grocery{}
	var err = db.DB.Where("id = ?", id).Find(&oldGrocery).Error
	if err != nil {
		return types.Grocery{}, err
	} else {
		err = db.DB.Where("id = ?", id).Delete(&oldGrocery).Error
		if err != nil {
			return types.Grocery{}, err
		}
		return oldGrocery, nil
	}
}
