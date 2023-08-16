package helper

import (
	"gotest/types"

	"gorm.io/gorm"
)

func CreateGrocery(db *gorm.DB, grocery *types.Grocery) (err error) {
	err = db.Create(grocery).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGroceryByName(db *gorm.DB, name string) (groceries []types.Grocery) {
	var groceriesList = []types.Grocery{}
	var err = db.Where("name = ?", name).Find(&groceriesList).Error
	if err != nil {
		return []types.Grocery{}
	}
	return groceriesList
}

func UpdateGrocery(db *gorm.DB, grocery types.Grocery) (updatedGrocery types.Grocery) {
	var newGrocery = types.Grocery{}
	var err = db.Save(&grocery).Where("id = ?", grocery.ID).Find(&newGrocery).Error
	if err != nil {
		return types.Grocery{}
	}
	return newGrocery
}

func DeleteGroceryById(db *gorm.DB, id int) (grocery types.Grocery) {
	var oldGrocery = types.Grocery{}
	var err = db.Where("id = ?", id).Find(&oldGrocery).Error
	if err != nil {
		return types.Grocery{}
	} else {
		err = db.Delete(&oldGrocery).Error
		if err != nil {
			return types.Grocery{}
		}
		return oldGrocery
	}
}

func GetAllGroceries(db *gorm.DB) (groceries []types.Grocery) {
	var groceryList []types.Grocery
	var err = db.Find(&groceryList).Error
	if err != nil {
		return []types.Grocery{}
	}
	return groceryList
}
