//go:generate mockery --name=GroceryServiceInterface
package service

import (
	"fmt"
	"gotest/db"
	"gotest/types"
	"strings"
)

type GroceryService struct {
	groceryRepo db.GroceryRepository
	userRepo    db.UserRepository
}

type GroceryServiceInterface interface {
	GetAllGroceries() (groceries []types.Grocery)
	SearchGroceriesFromUser(query string) (groceries []types.Grocery)
	GetAllGroceriesFromUser(userId string) (groceries []types.Grocery, err error)
	GetUserIdByUsername(username string) (userId string, err error)
	CreateGroceryForUser(username string, groceryId int) (err error)
	CreateGrocery(grocery *types.Grocery) (groceryNew types.Grocery, err error)
	DeleteGroceryForUser(username string, groceryId int) (err error)
	GetGroceryByName(name string) (groceries []types.Grocery)
	UpdateGrocery(grocery types.Grocery) (updatedGrocery types.Grocery, errStatus error)
	UpdateStatusOfGrocery(groceryId int, status bool) (newGrocery types.Grocery, err error)
	DeleteGroceryById(id int) (grocery types.Grocery, errStatus error)
}

func NewGroceryService(ds *db.MariaDBDataStore, groceryRepo db.GroceryRepository, userRepo db.UserRepository) *GroceryService {
	return &GroceryService{
		groceryRepo: groceryRepo,
		userRepo:    userRepo,
	}
}

func (gs *GroceryService) GetAllGroceries() (groceries []types.Grocery) {
	var groceryList []types.Grocery
	groceryList, err := gs.groceryRepo.GetAllGroceries()
	fmt.Printf("Hallo3 %v", groceryList)
	if err != nil {
		return []types.Grocery{}
	}
	return groceryList
}

func (gs *GroceryService) SearchGroceriesFromUser(query string) (groceriesList []types.Grocery) {
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

func (gs *GroceryService) GetAllGroceriesFromUser(userId string) (groceries []types.Grocery, err error) {
	userGroceries, err := gs.groceryRepo.GetAllGroceriesFromUser(userId)
	if err != nil {
		return []types.Grocery{}, err
	}
	for _, userGrocery := range userGroceries {
		grocery, err := gs.groceryRepo.FindGroceryWithId(userGrocery.GroceryID)
		if err != nil {
			return []types.Grocery{}, err
		}
		groceries = append(groceries, grocery)
	}
	return groceries, nil
}

func (gs *GroceryService) CreateGrocery(grocery *types.Grocery) (groceryNew types.Grocery, err error) {
	newGrocery, err := gs.groceryRepo.CreateGrocery(*grocery)
	if err != nil {
		return types.Grocery{}, err
	}
	return newGrocery, err
}

func (gs *GroceryService) CreateGroceryForUser(username string, groceryId int) (err error) {
	user, err := gs.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	err = gs.groceryRepo.CreateGroceryForUser(user.ID, groceryId)
	return err
}

func (gs *GroceryService) DeleteGroceryForUser(username string, groceryId int) (err error) {
	deletedGrocery, err := gs.DeleteGroceryById(groceryId)
	if err != nil {
		return
	}
	user, err := gs.userRepo.GetUserByUsername(username)
	if err != nil {
		return
	}
	err = gs.groceryRepo.DeleteGroceryFromUser(user.ID, int(deletedGrocery.ID))
	return err
}

func (gs *GroceryService) GetGroceryByName(name string) (groceries []types.Grocery) {
	groceriesList, err := gs.groceryRepo.GetGroceriesByName(name)
	if err != nil {
		return []types.Grocery{}
	}
	return groceriesList
}

func (gs *GroceryService) UpdateGrocery(grocery types.Grocery) (updatedGrocery types.Grocery, errStatus error) {
	updatedGrocery, err := gs.groceryRepo.UpdateGrocery(grocery)
	return updatedGrocery, err
}

func (gs *GroceryService) UpdateStatusOfGrocery(groceryId int, status bool) (newGrocery types.Grocery, err error) {
	grocery, err := gs.groceryRepo.UpdateStatusOfGrocery(groceryId, status)
	return grocery, err
}

func (gs *GroceryService) DeleteGroceryById(id int) (grocery types.Grocery, errStatus error) {
	oldGrocery, err := gs.groceryRepo.DeleteGrocery(id)
	return oldGrocery, err
}
