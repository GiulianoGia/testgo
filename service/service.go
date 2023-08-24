package service

import "gotest/db"

type Container struct {
	GroceryService GroceryService
	UserService    UserService
}

func NewContainer(ds *db.MariaDBDataStore, groceryRepo db.GroceryRepository, userRepository db.UserRepository) Container {
	return Container{
		GroceryService: *NewGroceryService(ds, groceryRepo, userRepository),
		UserService:    *NewUserService(ds, userRepository),
	}
}
