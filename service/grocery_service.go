package service

import (
	"gotest/types"
	"strings"
)

func (gs *ServiceStruct) GetAllGroceries() (groceries []types.Grocery) {
	var groceryList []types.Grocery
	groceryList, err := gs.dataStore.GetAllGroceries()
	if err != nil {
		return []types.Grocery{}
	}
	return groceryList
}

func (gs *ServiceStruct) SearchGroceriesFromUser(query string) (groceriesList []types.Grocery) {
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

func (gs *ServiceStruct) GetAllGroceriesFromUser(userId string) (groceries []types.Grocery, err error) {
	userGroceries, err := gs.dataStore.GetAllGroceriesFromUser(userId)
	if err != nil {
		return []types.Grocery{}, err
	}
	for _, userGrocery := range userGroceries {
		grocery, err := gs.dataStore.FindGroceryWithId(userGrocery.GroceryID)
		if err != nil {
			return []types.Grocery{}, err
		}
		groceries = append(groceries, grocery)
	}
	return groceries, nil
}

func (gs *ServiceStruct) CreateGrocery(grocery *types.Grocery) (groceryNew types.Grocery, err error) {
	newGrocery, err := gs.dataStore.CreateGrocery(*grocery)
	if err != nil {
		return types.Grocery{}, err
	}
	return newGrocery, err
}

func (gs *ServiceStruct) CreateGroceryForUser(username string, groceryId int) (err error) {
	user, err := gs.dataStore.GetUserByUsername(username)
	if err != nil {
		return err
	}
	err = gs.dataStore.CreateGroceryForUser(user.ID, groceryId)
	return err
}

func (gs *ServiceStruct) DeleteGroceryForUser(username string, groceryId int) (err error) {
	user, err := gs.dataStore.GetUserByUsername(username)
	if err != nil {
		return
	}
	err = gs.dataStore.DeleteGroceryFromUser(user.ID, groceryId)
	if err != nil {
		return
	}
	_, err = gs.DeleteGroceryById(groceryId)
	return err
}

func (gs *ServiceStruct) GetGroceryByName(name string) (groceries []types.Grocery) {
	groceriesList, err := gs.dataStore.GetGroceriesByName(name)
	if err != nil {
		return []types.Grocery{}
	}
	return groceriesList
}

func (gs *ServiceStruct) UpdateGrocery(grocery types.Grocery) (updatedGrocery types.Grocery, errStatus error) {
	updatedGrocery, err := gs.dataStore.UpdateGrocery(grocery)
	return updatedGrocery, err
}

func (gs *ServiceStruct) UpdateStatusOfGrocery(groceryId int, status bool) (newGrocery types.Grocery, err error) {
	grocery, err := gs.dataStore.UpdateStatusOfGrocery(groceryId, status)
	return grocery, err
}

func (gs *ServiceStruct) DeleteGroceryById(id int) (grocery types.Grocery, errStatus error) {
	oldGrocery, err := gs.dataStore.DeleteGrocery(id)
	return oldGrocery, err
}
